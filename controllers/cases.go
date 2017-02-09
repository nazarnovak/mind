package controllers

import(
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/nazarnovak/mind/data"
	"github.com/gorilla/mux"
)

func NewCase(w http.ResponseWriter, r *http.Request) {
	user, err := getSessionUser(r)
	if err != nil || user == nil {
		log.Println("Couldn't get user from session when creating " +
			"a case")
		serveBadRequest(w, r)
		return
	}

	c := data.Case{}
	c.CreatorId = user.ID

	id, err := c.Put()
	if err != nil {
		log.Println("Error while creating new case")
		serveBadRequest(w, r)
		return
	}

	err = data.GreetMessage(id)
	if err != nil {
		log.Println("Error while creating greet message")
		serveInternalServerError(w, r)
		return
	}

	http.Redirect(w, r, "/cases/" + strconv.Itoa(id), 303)
}

func GetCase(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	caseIdStr := vars["caseId"]

	c, err := data.GetCaseById(caseIdStr)
	if err != nil {
		log.Println("Couldn't get case id " + caseIdStr)
		serveInternalServerError(w, r)
		return
	}
	if c == nil {
		log.Println("Cased id " + caseIdStr + " not found in db")
		serveNotFound(w, r)
		return
	}

	user, err:= getSessionUser(r)
	if err != nil || user == nil {
		log.Println("Couldn't get user from session when creating " +
			"a case")
		serveBadRequest(w, r)
		return
	}

	data := struct{
		CaseId string
		UserId string
		UserName string
	}{
		caseIdStr,
		strconv.Itoa(user.ID),
		user.Name,
	}

	tpl := template.Must(template.ParseFiles(
	"public/templates/layout.gohtml", "public/templates/case.gohtml"))

	err = tpl.Execute(w, data)
	if err != nil {
		log.Println("Error while loding case template")
		serveInternalServerError(w, r)
		return
	}
}

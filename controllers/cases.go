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
	if err != nil {
		log.Println("Error while getting user from session: " + err.Error())
		serveInternalServerError(w, r)
		return
	}

	if user == nil {
		log.Println("Not logged in")
		serveNotFound(w, r)
		return
	}

	if user.Role != data.ROLEPATIENT {
		log.Println("Only patients can create cases")
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
	var c *data.Case

	vars := mux.Vars(r)
	caseIdStr := vars["caseId"]
	caseId, err := strconv.Atoi(caseIdStr)

	user, err := getSessionUser(r)
	if err != nil {
		log.Println("Error while getting user from session: " + err.Error())
		serveInternalServerError(w, r)
		return
	}

	if user == nil {
		log.Println("Not logged in")
		serveNotFound(w, r)
		return
	}

	c, err = data.GetCaseById(caseId)

	if user.Role != data.ROLEDOCTOR && user.ID != c.CreatorId {
		log.Println("Patient can't view cases of other " +
			"patients")
		serveNotFound(w, r)
		return
	}

	if err != nil {
		log.Println("Couldn't get case id " + caseIdStr + ":" + err.Error())
		serveInternalServerError(w, r)
		return
	}
	if c == nil {
		log.Println("Case id " + caseIdStr + " not found in db; " +
			"User id " + strconv.Itoa(user.ID))
		serveNotFound(w, r)
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

package controllers

import(
	"html/template"
	"log"
	"net/http"
	//"github.com/gorilla/sessions"
	"github.com/nazarnovak/mind/data"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var userName string

	user, err := getSessionUser(r)
	if err != nil {
		log.Println("Error while getting user from session: " + err.Error())
		serveInternalServerError(w, r)
		return
	}

	if user != nil {
		userName = user.Name
	}

	casesIds, err := data.GetCasesByCreatorId(user.ID)

	data := struct{
		User string
		Cases []int
	}{
		userName,
		casesIds,
	}

	tpl := template.Must(template.ParseFiles(
	"public/templates/layout.gohtml", "public/templates/index.gohtml"))
	tpl.Execute(w, data)
}
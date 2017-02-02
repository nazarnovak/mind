package public

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nazarnovak/mind/data"
)

var TplIndex = template.Must(template.ParseFiles(
	"public/templates/layout.gohtml", "public/templates/index.gohtml"))
var TplCase = template.Must(template.ParseFiles(
	"public/templates/layout.gohtml", "public/templates/case.gohtml"))

func Home(w http.ResponseWriter, r *http.Request) {
	// Load user cases and pass to template, to show user's existing ones
	TplIndex.Execute(w, nil)
}

func GetCase(w http.ResponseWriter, r *http.Request) {
// To get from cookie
UserId := "111"
UserName := "Nazar Novak"
	vars := mux.Vars(r)
	caseIdStr := vars["caseId"]

	c, err := data.GetCaseById(caseIdStr)
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
	if c == nil {
		ServeNotFound(w, r)
		return
	}

	data := struct{
		CaseId string
		UserId string
		UserName string
	}{caseIdStr, UserId, UserName}

	err = TplCase.Execute(w, data)
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
}

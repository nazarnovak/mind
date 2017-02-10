package controllers

import(
	"html/template"
	"log"
	"net/http"

	"github.com/nazarnovak/mind/data"
	"errors"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var userName string
	var userRole int
	var casesIds []int

	user, err := getSessionUser(r)
	if err != nil {
		log.Println("Error while getting user from session: " + err.Error())
		serveInternalServerError(w, r)
		return
	}

	if user != nil {
		userName = user.Name
		userRole = user.Role

		switch user.Role {
		case data.ROLEPATIENT:
			casesIds, err = data.GetCasesByCreatorId(user.ID)
		case data.ROLEDOCTOR:
			casesIds, err = data.GetCasesByDoctorId(user.ID)
		default:
			err = errors.New("Unknown role")
		}

		if err != nil {
			log.Println(err)
			serveInternalServerError(w, r)
			return
		}
	}

	data := struct{
		UserName string
		UserRole int
		Cases []int
	}{
		userName,
		userRole,
		casesIds,
	}

	tpl := template.Must(template.ParseFiles(
	"public/templates/layout.gohtml", "public/templates/index.gohtml"))
	tpl.Execute(w, data)
}
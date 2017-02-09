package controllers

import(
	"net/http"
	"log"

	"github.com/nazarnovak/mind/data"
)



func Login(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("user")

	user, err := data.GetUserByName(userName)
	if err != nil {
		log.Println("Error when getting user: " + err.Error())
		serveInternalServerError(w, r)
		return
	}

	if user == nil {
		log.Println("Incorrect user on login: " + userName)
		serveBadRequest(w, r)
		return
	}

	err = setSessionUser(w, r, user)
	if err != nil {
		log.Println("Error when setting user into session: " +
			err.Error())
		serveInternalServerError(w, r)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	deleteSessionUser(w, r)

	http.Redirect(w, r, "/", http.StatusFound)
}

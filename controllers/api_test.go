package controllers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/nazarnovak/mind/data"
)

var baseUrl = "http://localhost:8080"
//var userName = "Nazar Novak"
var user = data.User{111, "Nazar Novak", 0}
var caseId = 15

// CreateCaseMessage
func ATestNoUserInSession(t *testing.T) {
	caseIdStr := strconv.Itoa(caseId)

	req, err := http.NewRequest("POST", baseUrl + "/cases/" + caseIdStr +
		"/events",nil)
	if err != nil {
		log.Fatalln(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCaseMessage)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Handler returned wrong code: expected %v, got %v",
			http.StatusNotFound, rr.Code)
	}
}

//func loginUser() {
//	loginUrl := baseUrl + "/login"
//
//	_, err := http.PostForm(loginUrl, url.Values{"user": {userName}})
//	if err != nil {
//		log.Fatalln(err)
//	}
//}

func TestNoCaseFound(t *testing.T) {
	//loginUser()
	caseIdStr := strconv.Itoa(caseId)

	reqUrl := baseUrl + "/api/cases/" + caseIdStr + "/events"

	req, err := http.NewRequest("POST", reqUrl,nil)
	if err != nil {
		log.Fatalln(err)
	}

	rr := httptest.NewRecorder()

	Router := mux.NewRouter()
	Router.NewRoute().
		Methods("POST").
		Path("/api/cases/{caseId:[0-9]+}/events").
		HandlerFunc(CreateCaseMessage)
// set in Gorilla session
	setSessionUser(rr, req, &user)
	Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Handler returned wrong code: expected %v, got %v",
			http.StatusNotFound, rr.Code)
	}
}

// CreateCaseMessage
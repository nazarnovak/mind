package controllers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/nazarnovak/mind/data"
	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

var baseUrl = "http://localhost:8080/"
var user = data.User{111, "Nazar", 0}
var caseId = 15

// CreateCaseMessage
func TestNoUserInSession(t *testing.T) {
	caseIdStr := strconv.Itoa(caseId)

	req, err := http.NewRequest("POST", "/cases/" + caseIdStr + "/events",nil)
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

func ATestNoCaseFound(t *testing.T) {
	caseIdStr := strconv.Itoa(caseId)

	reqUrl := "/api/cases/" + caseIdStr + "/events"
log.Println(reqUrl)
	req, err := http.NewRequest("POST", reqUrl,nil)
	if err != nil {
		log.Fatalln(err)
	}

	rr := httptest.NewRecorder()
//Extract caseId from url/use Gorilla mux
	route := Router.NewRoute().
		Methods("POST").
		Path("/api/cases/{caseId:[0-9]+}/events").
		HandlerFunc(CreateCaseMessage)

	setSessionUser(rr, req, &user)

	handler := route.GetHandler()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Handler returned wrong code: expected %v, got %v",
			http.StatusNotFound, rr.Code)
	}
}

// CreateCaseMessage
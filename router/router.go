package router

import (
	"github.com/gorilla/mux"

	"github.com/nazarnovak/mind/api"
	"github.com/nazarnovak/mind/controllers"
)

var Router = mux.NewRouter()

func SetRoutes() {
	Router.NewRoute().
		Methods("GET").
		Path("/").
		HandlerFunc(controllers.Index)
	Router.NewRoute().
		Methods("POST").
		Path("/login").
		HandlerFunc(controllers.Login)
	Router.NewRoute().
		Methods("GET").
		Path("/logout").
		HandlerFunc(controllers.Logout)
	Router.NewRoute().
		Methods("GET").
		Path("/cases/{caseId:[0-9]+}").
		HandlerFunc(controllers.GetCase)
	Router.NewRoute().
		Methods("POST").
		Path("/cases/new").
		HandlerFunc(controllers.NewCase)
	Router.NewRoute().
		Methods("GET").
		Path("/api/cases/{caseId:[0-9]+}/events").
		HandlerFunc(api.GetCaseMessages)
	Router.NewRoute().
		Methods("POST").
		Path("/api/cases/{caseId:[0-9]+}/events").
		HandlerFunc(api.CreateCaseMessage)
}
package main

import (
	"github.com/gorilla/mux"
	"github.com/nazarnovak/mind/public"
	"github.com/nazarnovak/mind/api"
)

var UIRouter = mux.NewRouter()

func init() {
	UIRouter.NewRoute().
		Methods("GET").
		Path("/").
		HandlerFunc(public.Home)
	UIRouter.NewRoute().
		Methods("GET").
		Path("/cases/{caseId:[0-9]+}").
		HandlerFunc(public.GetCase)
	UIRouter.NewRoute().
		Methods("POST").
		Path("/cases/new").
		HandlerFunc(api.NewCase)
	UIRouter.NewRoute().
		Methods("GET").
		Path("/api/cases/{caseId:[0-9]+}/events").
		HandlerFunc(api.GetCaseMessages)
	UIRouter.NewRoute().
		Methods("POST").
		Path("/api/cases/{caseId:[0-9]+}/events").
		HandlerFunc(api.CreateCaseMessage)
}

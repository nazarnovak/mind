package controllers

import (
	"net/http"
)

func serveBadRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Bad Request", http.StatusBadRequest)
}

func serveNotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Found", http.StatusNotFound)
}

func serveInternalServerError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

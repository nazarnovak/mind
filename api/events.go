package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/nazarnovak/mind/data"
)

const (
	supportCmd = "@support "
)

type Event struct {
	UserId  int `json:"userid,string"`
	UserName string `json:"username"`
	Message string `json:"msg"`

}

type Message struct {
	Id       string
	UserName string
	UserRole string
	Created  time.Time
	Content  string
}

func GetCaseMessages(w http.ResponseWriter, r *http.Request) {
	var msgs []data.Message

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

	sinceStr := r.URL.Query().Get("since")

	if sinceStr == "" {
		sinceStr = "0001-01-01T01:01:01+01:00"
	}
	_, err = time.Parse(time.RFC3339, sinceStr)

	if err != nil {
		ServeBadRequest(w, r)
		return
	}

	msgs, err = data.GetCaseEventsByCaseId(caseIdStr, sinceStr)
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(msgs)
	if err != nil {
		ServeInternalServerError(w, r)
	}
}

func CreateCaseMessage(w http.ResponseWriter, r *http.Request) {
	var c *data.Case
	var err error

	vars := mux.Vars(r)
	caseIdStr := vars["caseId"]

	c, err = data.GetCaseById(caseIdStr)
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
	if c == nil {
		ServeNotFound(w, r)
		return
	}

	e := Event{}
	err = json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		ServeBadRequest(w, r)
		return
	}

	if strings.HasPrefix(e.Message, supportCmd) {
		notifySupport(strings.TrimLeft(e.Message, supportCmd))
	}

	ce := data.CaseEvent{}
	ce.CaseId = c.ID
	ce.UserId = e.UserId
	ce.TypeId = data.Messages
	ce.Created = time.Now().Format(time.RFC3339)
	ce.Content = e.Message
	id, err := ce.Put()
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}

	err = data.Emit("case:" + strconv.Itoa(c.ID), "message:"+ strconv.Itoa(id))
	if err != nil {
		log.Println(err)
	}
}

func notifySupport(msg string) {
	log.Println("Support was notified about: " + msg)
}

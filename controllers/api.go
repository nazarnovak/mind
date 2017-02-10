package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/nazarnovak/mind/data"
	"errors"
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
	var c *data.Case
	var msgs []data.Message

	vars := mux.Vars(r)
	caseIdStr := vars["caseId"]
	caseId, err := strconv.Atoi(caseIdStr)
	if err != nil {
		log.Println("Error when converting case id to int: " + err.Error())
		serveInternalServerError(w, r)
		return
	}

	user, err := getSessionUser(r)
	if err != nil {
		log.Println("Error while getting user from session: " + err.Error())
		serveInternalServerError(w, r)
		return
	}

	if user == nil {
		log.Println("Not logged in")
		serveNotFound(w, r)
		return
	}

	switch user.Role {
	case data.ROLEPATIENT:
		c, err = data.GetCaseByIdCreatorId(caseId, user.ID)
	case data.ROLEDOCTOR:
		c, err = data.GetCaseByIdDoctorId(caseId, user.ID)
	default:
		err = errors.New("Incorrect role id")
	}

	if err != nil {
		log.Println("Error while getting case with id " + caseIdStr +
			":" + err.Error())
		serveInternalServerError(w, r)
		return
	}
	if c == nil {
		log.Println("Error finding case with id " + caseIdStr)
		serveNotFound(w, r)
		return
	}

	sinceStr := r.URL.Query().Get("since")

	if sinceStr == "" {
		sinceStr = "0001-01-01T01:01:01+01:00"
	}
	_, err = time.Parse(time.RFC3339, sinceStr)

	if err != nil {
		serveBadRequest(w, r)
		return
	}

	msgs, err = data.GetCaseEventsByCaseId(caseIdStr, sinceStr)
	if err != nil {
		serveInternalServerError(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(msgs)
	if err != nil {
		serveInternalServerError(w, r)
	}
}

func CreateCaseMessage(w http.ResponseWriter, r *http.Request) {
	var c *data.Case

	vars := mux.Vars(r)
	caseIdStr := vars["caseId"]
	caseId, err := strconv.Atoi(caseIdStr)
	if err != nil {
		log.Println("Error when converting case id to int: " + err.Error())
		serveInternalServerError(w, r)
		return
	}

	user, err := getSessionUser(r)
	if err != nil {
		log.Println("Error while getting user from session: " + err.Error())
		serveInternalServerError(w, r)
		return
	}

	if user == nil {
		log.Println("Not logged in")
		serveNotFound(w, r)
		return
	}

	switch user.Role {
	case data.ROLEPATIENT:
		c, err = data.GetCaseByIdCreatorId(caseId, user.ID)
	case data.ROLEDOCTOR:
		c, err = data.GetCaseByIdDoctorId(caseId, user.ID)
	default:
		err = errors.New("Incorrect role id")
	}
	if err != nil {
		log.Println("Error while getting case with id " + caseIdStr +
			":" + err.Error())
		serveInternalServerError(w, r)
		return
	}
	if c == nil {
		log.Println("Error finding case with id " + caseIdStr)
		serveNotFound(w, r)
		return
	}

	e := Event{}
	err = json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		log.Println("Error when decoding AJAX request:" + err.Error())
		serveBadRequest(w, r)
		return
	}

	err = checkUserMessage(c.ID, c.DoctorId, e.Message)
	if err != nil {
		log.Println("Error checking user message:" + err.Error())
	}

	ce := data.CaseEvent{}
	ce.CaseId = c.ID
	ce.UserId = e.UserId
	ce.TypeId = data.Messages
	ce.Created = time.Now().Format(time.RFC3339)
	ce.Content = e.Message
	id, err := ce.Put()
	if err != nil {
		log.Println("Error when creating a case event: " + err.Error())
		serveInternalServerError(w, r)
		return
	}

	err = data.Emit("case:" + strconv.Itoa(c.ID), "message:"+ strconv.Itoa(id))
	if err != nil {
		log.Println(err)
		return
	}
}

func checkUserMessage(caseId int, doctorId int, msg string) error {
	if msg == "/submit" && doctorId == 0 {
		err := assignDoctor(caseId)
		if err != nil {
			return err
		}
	}

	if strings.HasPrefix(msg, supportCmd) {
		var sMsg= strings.TrimLeft(msg, supportCmd)
		notifySupport(sMsg)
		return nil
	}

	return nil
}

func assignDoctor(caseId int) error {
	doctorId, err := data.GetRandomDoctorId()
	if err != nil {
		return err
	}

	if doctorId == 0 {
		return errors.New("There are no doctors in DB")
	}

	err = data.AssignDoctorToCase(doctorId, caseId)
	if err != nil {
		return err
	}

	return nil
}

func notifySupport(msg string) {
	log.Println("Support was notified about: " + msg)
}

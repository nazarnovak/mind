package controllers

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/nazarnovak/mind/data"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))
var sessionName = "session-name"

func getSessionUser(r *http.Request) (*data.User, error) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return nil, err
	}

	if(session.Values["userid"] == nil ||
	session.Values["username"] == nil ||
	session.Values["userrole"] == nil) {
		return nil, nil
	}

	user := &data.User{}
	user.ID = session.Values["userid"].(int)
	user.Name = session.Values["username"].(string)
	user.Role = session.Values["userrole"].(int)

	return user, nil
}

func setSessionUser(w http.ResponseWriter, r *http.Request, user *data.User) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}

	session.Values["userid"] = user.ID
	session.Values["username"] = user.Name
	session.Values["userrole"] = user.Role
	session.Save(r, w)

	return nil
}

func deleteSessionUser(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}

	delete(session.Values, "userid")
	delete(session.Values, "username")
	delete(session.Values, "userrole")
	session.Save(r, w)

	return nil
}
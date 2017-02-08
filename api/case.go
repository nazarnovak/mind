package api

import(
	"net/http"
	"strconv"

	"github.com/nazarnovak/mind/data"
)

func NewCase(w http.ResponseWriter, r *http.Request) {
	c := data.Case{}
	id, err := c.Put()
	if err != nil {
		ServeBadRequest(w, r)
		return
	}

	err = data.GreetMessage(id)
	if err != nil {
		ServeInternalServerError(w, r)
	}

	http.Redirect(w, r, "/cases/" + strconv.Itoa(id), 303)
}
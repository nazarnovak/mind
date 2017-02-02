package api

import(
	"encoding/json"
	"net/http"

	"github.com/nazarnovak/mind/data"
)

func NewCase(w http.ResponseWriter, r *http.Request) {
	c := data.Case{}
	id, err := c.Put()
	if err != nil {
		ServeBadRequest(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(id)
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}

	err = data.GreetMessage(id)
	if err != nil {
		ServeInternalServerError(w, r)
	}
}
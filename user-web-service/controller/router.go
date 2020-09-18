package controller

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers : set up all the routing required
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/", uc)
	http.Handle("/users", uc)
	http.Handle("/users/", uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

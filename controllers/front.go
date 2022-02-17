package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func RegisterControllers(logger logrus.FieldLogger) {
	uc := newUserController(logger)

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

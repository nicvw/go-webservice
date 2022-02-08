package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}

// constructor function
func newUserController() *userController {
	return &userController{ // this works because this is a struct data type
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

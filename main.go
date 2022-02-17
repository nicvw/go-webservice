package main

import (
	"net/http"

	"github.com/nicvw/go-webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		println(err.Error())
	}
}

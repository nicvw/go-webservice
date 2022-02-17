package main

import (
	"net/http"

	"github.com/nicvw/go-webservice/controllers"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{DisableTimestamp: true}
	controllers.RegisterControllers(logger)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		logger.Panic(err.Error())
	}
}

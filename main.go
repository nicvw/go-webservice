package main

import (
	"fmt"

	"github.com/nicvw/go-webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Bob",
		LastName:  "Burger",
	}
	fmt.Println(u)
}

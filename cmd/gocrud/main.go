package main

import (
	"fmt"
	"net/http"

	"github.com/andisis/gocrud/src/router"
	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
)

func main() {
	router := router.GetRouter()

	fmt.Println("Server is running...")

	err := http.ListenAndServe(":8000", handlers.CORS()(router))
	if err != nil {
		logrus.Error(err)
		fmt.Print(err)
	}
}

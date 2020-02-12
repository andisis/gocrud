package router

import (
	"net/http"
	"time"

	userHandler "github.com/andisis/gocrud/src/api/user/delivery/http"
	userRepo "github.com/andisis/gocrud/src/api/user/repository"
	userUsecase "github.com/andisis/gocrud/src/api/user/usecase"

	"github.com/andisis/gocrud/src/database"

	"github.com/gorilla/mux"
)

// GetRouter function to send the router to the main package
func GetRouter() *mux.Router {
	connection, _ := database.ConnectSQL()

	route := mux.NewRouter().StrictSlash(true)
	route.HandleFunc("/", indexHandler).Methods("GET")

	subrouter := route.PathPrefix("/api/v1").Subrouter()
	subrouter.HandleFunc("/", indexHandler).Methods("GET")
	timeoutContext := time.Duration(2) * time.Second

	// User handler
	usrRepo := userRepo.NewSQLUserRepository(connection.SQL)
	usrUsecase := userUsecase.NewUserUsecase(usrRepo, timeoutContext)
	userHandler.NewUserHandler(subrouter, usrUsecase)

	route.Handle("/", http.StripPrefix("/", subrouter))
	return route
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	html := "<div style=\"height: 100%;position: relative;\">" +
		"<div style=\"width: 100%;margin: 0;position: absolute;top: 50%;-ms-transform: translateY(-50%);transform: translateY(-50%);text-align: center;\">" +
		"<h3 style=\"font-family: Sans-serif;text-align: center;\">GOCRUD is ready...</h3>" +
		"<img src=\"https://camo.githubusercontent.com/c70f18274a81ee98dca1c116b68d5a35847b2e65/687474703a2f2f7374617469632e76656c76657463616368652e6f72672f70616765732f323031382f30362f31332f70617274792d676f706865722f64616e63696e672d676f706865722e676966\">" +
		"</div>" +
		"</div>"
	var message = html
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(message))
}

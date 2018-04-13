package main

import (
	"log"
	"net/http"
	"routers"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(false)
	// Routes for the User entity
	router = routers.SetBookingsRouters(router)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", router)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//return router
}

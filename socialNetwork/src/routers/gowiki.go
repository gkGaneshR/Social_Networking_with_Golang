package routers

import (
	"controllers"
	"fmt"
	//"log"
	"models/frontend"
	"net/http"

	"github.com/gorilla/mux"
)

func SetBookingsRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.IndexPageHandler)
	router.HandleFunc("/errorlogin", controllers.IndexPageHandler1)
	router.HandleFunc("/internal", controllers.InternalPageHandler)
	router.HandleFunc("/internalnew", controllers.InternalPageHandler1)
	router.HandleFunc("/register/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, frontend.RegisterPage)
	})
	router.HandleFunc("/newlogin", controllers.RegisterInDb).Methods("POST")
	router.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
	router.HandleFunc("/logout", controllers.LogoutHandler).Methods("POST")

	router.HandleFunc("/upload/", controllers.HelloServer)

	router.HandleFunc("/app/", controllers.Uploadimage)

	router.HandleFunc("/images/", controllers.HandleImages)
	//r.HandleFunc("/images/{ImageName}/", HandleImage)

	router.HandleFunc("/viewimage/", controllers.Viewimage)
	router.HandleFunc("/writecmnttodb/", controllers.WriteCmntToDb)
	router.HandleFunc("/readcmntfromdb/", controllers.ReadCmntFromDb)

	return router
}

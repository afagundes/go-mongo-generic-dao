package router

import (
	"github.com/afagundes/mongo-generic-dao/api/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/usuarios", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/usuario/{id}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/usuario/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/usuario/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/usuario", controllers.CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

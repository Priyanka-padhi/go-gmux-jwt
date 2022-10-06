package main

import (
	"github.com/gorilla/mux"
	"go-gmux-jwt/controllers/authcontroller"
	"go-gmux-jwt/models"
	"log"
	"net/http"
)

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	models.ConnectDB()
	initializeRouter()
}

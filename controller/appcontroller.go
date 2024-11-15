package controller

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func RouteInit() {
	headersOk := handlers.AllowedHeaders([]string{"Accept", "Authorization", "Content-Type", "Origin"})
	originsOk := handlers.AllowedOrigins([]string{
		"http://127.0.0.1:8000", // port Flutter web allowed
		"http://feliandra.my.id",
	})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})

	router := mux.NewRouter()

	userroute := router.PathPrefix("/users").Subrouter()
	userroute.HandleFunc("", GetUsers).Methods(http.MethodGet)
	userroute.HandleFunc("/add", AddUser).Methods(http.MethodPost)

	http.ListenAndServe(
		"127.0.0.1:8081",
		handlers.CORS(headersOk, originsOk, methodsOk)(router),
	)
}

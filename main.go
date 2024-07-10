package main

import (
	"fmt"
	"go_api_mux_gorm/handlers"
	"go_api_mux_gorm/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Starting Program...")


	//Enable database Migrations
	// This will create the tables in the database, better keep commet this line after the first execution
	//models.Migrations()

	// Create a new router
	mux := mux.NewRouter()

	prefix := "/api/v1/"

	// Group of routes
	mux.HandleFunc(prefix+"example", handlers.Example_get).Methods("GET")
	mux.HandleFunc(prefix+"querystring", handlers.Example_get_querystring).Methods("GET")
	mux.HandleFunc(prefix+"example/{id:[0-9]+}", handlers.Example_get_with_parameters).Methods("GET")
	mux.HandleFunc(prefix+"example", handlers.Example_post).Methods("POST")
	mux.HandleFunc(prefix+"example/{id:[0-9]+}", handlers.Example_put).Methods("PUT")
	mux.HandleFunc(prefix+"example/{id:[0-9]+}", handlers.Example_delete).Methods("DELETE")

	// Upload file
	mux.HandleFunc(prefix+"upload", handlers.Example_upload).Methods("POST")

	// CORS
	handler := cors.AllowAll().Handler(mux)

	//log.Fatal(http.ListenAndServe(":8084", mux))

	// Start the server usig the mux router and CORS
	log.Fatal(http.ListenAndServe(":8084", handler))

}

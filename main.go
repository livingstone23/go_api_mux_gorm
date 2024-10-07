package main

import (
	"fmt"
	"go_api_mux_gorm/handlers"
	"go_api_mux_gorm/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Starting Program...")

	//Enable database Migrations
	// This will create the tables in the database, better keep commet this line after the first execution
	//"go_api_mux_gorm/models"
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

	// Categories routes
	mux.HandleFunc(prefix+"categories", handlers.Category_get).Methods("GET")
	mux.HandleFunc(prefix+"categories/{id:[0-9]+}", handlers.Category_with_parameters).Methods("GET")
	mux.HandleFunc(prefix+"categories", handlers.Category_new).Methods("POST")
	mux.HandleFunc(prefix+"categories/{id:[0-9]+}", handlers.Category_update).Methods("PUT")
	mux.HandleFunc(prefix+"categories/{id:[0-9]+}", handlers.Category_delete).Methods("DELETE")

	// Products routes
	mux.HandleFunc(prefix+"products", handlers.Product_get).Methods("GET")
	mux.HandleFunc(prefix+"products", handlers.Product_new).Methods("POST")
	mux.HandleFunc(prefix+"products/{id:[0-9]+}", handlers.Product_with_parameters).Methods("GET")
	mux.HandleFunc(prefix+"products/{id:[0-9]+}", handlers.Product_update).Methods("PUT")
	mux.HandleFunc(prefix+"products/{id:[0-9]+}", handlers.Product_delete).Methods("DELETE")

	// Products Pictures routes
	mux.HandleFunc(prefix+"products-pictures/{id:[0-9]+}", handlers.ProductPicture_Upload).Methods("POST")
	mux.HandleFunc(prefix+"products-pictures/{id:[0-9]+}", handlers.ProductPicture_GetByProduct).Methods("GET")
	mux.HandleFunc(prefix+"products-pictures/{id:[0-9]+}", handlers.ProductPicture_Delete).Methods("DELETE")


	//Security routes
	mux.HandleFunc(prefix+"security/register", handlers.Security_register).Methods("POST")
	mux.HandleFunc(prefix+"security/login", handlers.Security_login).Methods("POST")


	// Route Protected, using the Middleware
	mux.HandleFunc(prefix+"security/protected", middleware.ValidateJWT(handlers.Security_protected)).Methods("GET")


	// CORS
	handler := cors.AllowAll().Handler(mux)

	//log.Fatal(http.ListenAndServe(":8084", mux))

	// Start the server usig the mux router and CORS
	log.Fatal(http.ListenAndServe(":8084", handler))

}

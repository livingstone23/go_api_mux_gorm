package handlers

import (
	"encoding/json"
	"fmt"
	"go_api_mux_gorm/dto"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"github.com/gorilla/mux"
)

//Package handlers is used to handle the routing of the application

type GenericResponse struct {
	Status  string
	Message string
}

// Example_get is a function that handles the GET request to the /api/v1/example endpoint
func Example_get(response http.ResponseWriter, request *http.Request) {

	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	// Set a custom header
	response.Header().Set("LivingstoneCano", "www.livingstonecano.com")
	// Create a JSON response
	output, _ := json.Marshal(GenericResponse{"success", "hello method Example_get"})
	fmt.Fprintln(response, string(output))

}

func Example_get_querystring(response http.ResponseWriter, request *http.Request) {

	fmt.Fprintln(response, "hello method querystring | id = ", request.URL.Query().Get("id"))
}

func Example_get_with_parameters(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintln(response, "hello method Example_get_with_parameters | id: ", vars["id"])
}

func Example_post(response http.ResponseWriter, request *http.Request) {

	var category dto.CategoryDto
	err := json.NewDecoder(request.Body).Decode(&category)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return

	}

	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	// Set a custom header
	response.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	answer := map[string]string{
		"status":      "success",
		"message":     "Post method Example_post whith map and json.NewEncoder",
		"Name":        category.Name,
		"Description": category.Description,
	}

	//Esta manera no es la mas rescomendada
	//response.WriteHeader(201)
	//Manera mas adecuada.
	response.WriteHeader(http.StatusCreated)
	// Create a JSON response
	json.NewEncoder(response).Encode(answer)
}

/*

func Example_post(response http.ResponseWriter, request *http.Request) {

	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	// Set a custom header
	response.Header().Set("LivingstoneCano", "www.livingstonecano.com")

	answer := map[string]string{
		"status":  "success",
		"message": "Post method Example_post whith map and json.NewEncoder",
	}

	//Esta manera no es la mas rescomendada
	//response.WriteHeader(201)
	//Manera mas adecuada.
	response.WriteHeader(http.StatusCreated)
	// Create a JSON response
	json.NewEncoder(response).Encode(answer)
}
*/

/*
func Example_post(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Post method Example_post")
}
*/

func Example_put(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Put method Example_put")
}

func Example_delete(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Delete method Example_delete")
}

func Example_upload(response http.ResponseWriter, request *http.Request) {

	file, handler, _ := request.FormFile("file")
	var extension = strings.Split(handler.Filename, ".")[1]
	time := strings.Split(time.Now().String(), " ")
	var fileName = string(time[4][6:14]) + "." + extension
	var filestring string = "public/uploads/files" + fileName

	// Create a new file
	f, err := os.OpenFile(filestring, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		http.Error(response, "Error uploading file ! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(response, "Error copying file ! "+err.Error(), http.StatusBadRequest)
		return
	}

	answer := map[string]string{
		"status":  "success",
		"message": "file saved successfully",
		"file":    fileName,
	}

	// Set the response header
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(answer)

}

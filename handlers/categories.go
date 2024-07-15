package handlers

import (
	"encoding/json"
	"go_api_mux_gorm/database"
	"go_api_mux_gorm/dto"
	"go_api_mux_gorm/models"
	"net/http"

	"github.com/gorilla/mux"
	"strconv"
)

func Category_get(response http.ResponseWriter, request *http.Request) {
	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	data := models.Categories{}

	//Link datasbase to the model
	//database.Database.Find(&data)

	//Link datasbase to the model and order by id desc
	database.Database.Order("id desc").Find(&data)


	//Set the response status
	response.WriteHeader(http.StatusOK)

	//Return the data in JSON format
	json.NewEncoder(response).Encode(data)

}


func Category_with_parameters(response http.ResponseWriter, request *http.Request) {
	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	// Get the parameters from the request
	vars := mux.Vars(request)

	// Convert the id to integer
	//id, _ = strconv.Atoi(vars["id"])

	data := models.Category{}

	if err := database.Database.First(&data, vars["id"]).Error; err != nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode("Category not found")	
		return
	} else {
		//Set the response status
		response.WriteHeader(http.StatusOK)

		//Return the data in JSON format
		json.NewEncoder(response).Encode(data)
	}
}


func Category_new(response http.ResponseWriter, request *http.Request) {
	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	// Get the parameters from the request
	var category dto.CategoryDto

	if err:= json.NewDecoder(request.Body).Decode(&category); err != nil {
		anwser := map[string] string{
			"status": "error",
			"message": "Invalid request",
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(anwser)
		return
	}

	// Create a category to insert in the database
	data := models.Category{ Name: category.Name }

	database.Database.Create(&data)

	anwser := map[string] string{
		"status": "Ok",
		"message": "Created successfully",
	}

	//Set the response status
	response.WriteHeader(http.StatusCreated)
	//Return the data in JSON format
	json.NewEncoder(response).Encode(anwser)

}

func Category_update(response http.ResponseWriter, request *http.Request) {
	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	// Get the parameters from the request
	vars := mux.Vars(request)

	// Convert the id to integer
	id, _ := strconv.Atoi(vars["id"])

	// Get the parameters from the request
	var category dto.CategoryDto

	if err:= json.NewDecoder(request.Body).Decode(&category); err != nil {
		anwser := map[string] string{
			"status": "error",
			"message": "Error in the request",
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(anwser)
		return
	}

	// Create a category to insert in the database
	data := models.Category{}

	if err := database.Database.First(&data, id).Error; err != nil {
		anwser := map[string] string{
			"status": "error",
			"message": "Resource not found",
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(anwser)
		return
	} else {
		data.Name = category.Name
		database.Database.Save(&data)

		//Set the response status
		anwser := map[string] string{
			"status": "Ok",
			"message": "Created successfully",
		}
	
		//Set the response status
		response.WriteHeader(http.StatusCreated)
		//Return the data in JSON format
		json.NewEncoder(response).Encode(anwser)
	}


}


func Category_delete(response http.ResponseWriter, request *http.Request) {
	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	// Get the parameters from the request
	vars := mux.Vars(request)

	// Convert the id to integer
	id, _ := strconv.Atoi(vars["id"])


	// Create a category to insert in the database
	data := models.Category{}

	if err := database.Database.First(&data, id).Error; err != nil {
		anwser := map[string] string{
			"status": "error",
			"message": "Resource not found",
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(anwser)
		return
	} else {
		database.Database.Delete(&data)

		//Set the response status
		anwser := map[string] string{
			"status": "Ok",
			"message": "Register deleted successfully",
		}
	
		//Set the response status
		response.WriteHeader(http.StatusOK)
		//Return the data in JSON format
		json.NewEncoder(response).Encode(anwser)
	}


}
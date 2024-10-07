package handlers

import (
	"encoding/json"
	"go_api_mux_gorm/database"
	"go_api_mux_gorm/dto"
	"go_api_mux_gorm/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)


func Product_get(response http.ResponseWriter, request *http.Request) {
	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	data := models.Products{}

	//Link datasbase to the model
	//database.Database.Find(&data)

	//Link datasbase to the model and order by id desc
	//Preload is used to bring the relationship data
	database.Database.Order("id desc").Preload("Category").Find(&data)

	//Set the response status
	response.WriteHeader(http.StatusOK)

	//Return the data in JSON format
	json.NewEncoder(response).Encode(data)
}


func Product_new(response http.ResponseWriter, request *http.Request) {
	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	// Get the parameters from the request
	var product dto.ProductDto

	if err := json.NewDecoder(request.Body).Decode(&product); err != nil {
		anwser := map[string]string{
			"status":  "error",
			"message": "Invalid request",
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(anwser)
		return
	}

	// Create a category to insert in the database
	data := models.Product{Name: product.Name, Price: product.Price, Stock: product.Stock, Description: product.Description, CategoryID: product.CategoryID, DateRegister: time.Now()}

	database.Database.Create(&data)

	anwser := map[string]string{
		"status":  "Ok",
		"message": "Created successfully",
	}

	//Set the response status
	response.WriteHeader(http.StatusCreated)
	//Return the data in JSON format
	json.NewEncoder(response).Encode(anwser)
}


func Product_with_parameters(response http.ResponseWriter, request *http.Request) {
	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	// Get the parameters from the request
	vars := mux.Vars(request)

	// Convert the id to integer
	//id, _ = strconv.Atoi(vars["id"])

	data := models.Product{}

	if err := database.Database.Preload("Category").First(&data, vars["id"]).Error; err != nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode("Product not found")
		return
	} else {
		//Set the response status
		response.WriteHeader(http.StatusOK)

		//Return the data in JSON format
		json.NewEncoder(response).Encode(data)
	}
}


func Product_update(response http.ResponseWriter, request *http.Request) {
	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	// Get the parameters from the request
	vars := mux.Vars(request)

	// Convert the id to integer
	id, _ := strconv.Atoi(vars["id"])

	// Get the parameters from the request
	var product dto.ProductDto

	if err := json.NewDecoder(request.Body).Decode(&product); err != nil {
		anwser := map[string]string{
			"status":  "error",
			"message": "Error in the request",
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(anwser)
		return
	}

	// Create a category to insert in the database
	data := models.Product{}

	if err := database.Database.First(&data, id).Error; err != nil {
		anwser := map[string]string{
			"status":  "error",
			"message": "Resource not found",
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(anwser)
		return
	} else {

		data.Name = product.Name
		data.Price = product.Price
		data.Stock = product.Stock
		data.Description = product.Description
		data.CategoryID = product.CategoryID
		data.DateRegister = time.Now()

		database.Database.Save(&data)

		//Set the response status
		answer := map[string]string{
			"status":  "Ok",
			"message": "Created successfully",
		}

		//Set the response status
		response.WriteHeader(http.StatusCreated)
		//Return the data in JSON format
		json.NewEncoder(response).Encode(answer)
	}
}


func Product_delete(response http.ResponseWriter, request *http.Request) {
	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	// Get the parameters from the request
	vars := mux.Vars(request)

	// Convert the id to integer
	id, _ := strconv.Atoi(vars["id"])

	// Create a category to insert in the database
	data := models.Product{}

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
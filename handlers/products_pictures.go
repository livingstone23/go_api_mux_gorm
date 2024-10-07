package handlers

import (
	"encoding/json"
	"go_api_mux_gorm/database"
	"go_api_mux_gorm/models"
	"net/http"
	"os"
	"strings"
	"time"
	"io"
	"github.com/gorilla/mux"
	"strconv"
)



func ProductPicture_Upload(response http.ResponseWriter, request *http.Request) {
	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	//Set the response status
	response.WriteHeader(http.StatusOK)

	//Return the data in JSON format

	file, handler, _ := request.FormFile("file")
	var extension = strings.Split(handler.Filename, ".")[1]

	// Get the current time and format it
    currentTime := time.Now().Format("20060102150405")

	//picture := string(time[4][6:14]) + "." + extension

	// Create the picture name with the extension
    picture := currentTime + "." + extension

	var filename string ="public/uploads/products/" + picture
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		http.Error(response, "Error uploading the file !"+ err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(response, "Error uploading the file !"+ err.Error(), http.StatusBadRequest)
		return
	}
	
	//Create the register in the database
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	datos := models.ProductPicture{Name: picture, ProductID: uint(id)}
	database.Database.Create(&datos)


	answer := map[string]string{	
		"status": "Ok",
		"message": "File uploaded",
		"picture": picture,
	}

	json.NewEncoder(response).Encode(answer)
}

//Method for Get all the Pictures by Products
func ProductPicture_GetByProduct(response http.ResponseWriter, request *http.Request) {
	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	// Get the parameters from the request
	vars := mux.Vars(request)

	// Convert the id to integer
	id, _ := strconv.Atoi(vars["id"])

	product := models.Product{}

	if err := database.Database.First(&product, id).Error; err != nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode("Product not found")
		return
	} else {
		pictures := models.ProductPictures{}
		database.Database.Where("product_id = ?", id).Find(&pictures)
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(pictures)
	}
}


//Method for Delete the Picture by Id
func ProductPicture_Delete(response http.ResponseWriter, request *http.Request) {
	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	// Get the parameters from the request
	vars := mux.Vars(request)

	// Convert the id to integer
	id, _ := strconv.Atoi(vars["id"])

	// Create a category to insert in the database
	data := models.ProductPicture{}

	if err := database.Database.First(&data, id).Error; err != nil {
		anwser := map[string] string{
			"status": "error",
			"message": "Resource not found",
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(anwser)
		return
	} else {
		//Delete the file
		var filename string ="public/uploads/products/" + data.Name
		err := os.Remove(filename)
		if err != nil {
			anwser := map[string] string{
				"status": "error",
				"message": "Error deleting the file",

			}
			response.WriteHeader(http.StatusNotFound)
			json.NewEncoder(response).Encode(anwser)
			return
		}


		//Delete the register in the database
		database.Database.Delete(&data)

		//Set the response status
		anwser := map[string] string{
			"status": "Ok",
			"message": "Register deleted successfully",
		}
	
		//Set the response status
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(anwser)
	}
}
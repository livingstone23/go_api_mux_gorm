package handlers

import (
	"encoding/json"
	"go_api_mux_gorm/database"
	"go_api_mux_gorm/models"
	"net/http"
)

func Category_get(response http.ResponseWriter, request *http.Request) {
	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	data := models.Categories{}

	//Link datasbase to the model
	database.Database.Find(&data)

	response.WriteHeader(http.StatusOK)

	//Return the data in JSON format
	json.NewEncoder(response).Encode(data)

}

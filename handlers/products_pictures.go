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
	//response.Write([]byte(`{"message": "ProductPicture_Upload"}`))


	file, handler, _ := request.FormFile("file")
	var extension = strings.Split(handler.Filename, ".")[1]
	time := strings.Split(time.Now().String(), " ")
	picture := string(time[4][6:14]) + "." + extension
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
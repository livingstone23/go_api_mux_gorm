package handlers

import (
	"encoding/json"
	"go_api_mux_gorm/database"
	"go_api_mux_gorm/dto"
	"go_api_mux_gorm/models"
	"net/http"
	"time"
	"golang.org/x/crypto/bcrypt"
	"go_api_mux_gorm/jwt"

)


// Function to register a new user
func Security_register(response http.ResponseWriter, request *http.Request) {
	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	// Get the parameters from the request
	var user dto.UserDto

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		anwser := map[string]string{
			"status":  "error",
			"message": "Invalid request",
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(anwser)
		return
	}

	// Validate if exist record with the same email
	userDatabase := models.User{}
	if database.Database.Where("email = ?", user.Email).Find(&userDatabase).RowsAffected >0 {
		anwser := map[string]string{
			"status":  "error",
			"message": "User already registered with this email" + user.Email,
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(anwser)
		return
	} else 
	{
		//Generate the password hash with bcrypt
		costo := 8
		bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), costo)
		datos := models.User{Name: user.Name, Email: user.Email, Password: string(bytes), PerfilID: user.PerfilID, DateRegister: time.Now()}
		database.Database.Create(&datos)
		anwser := map[string]string{
			"status":  "Ok",
			"message": "User registered",
		}

		//Set the response status
		response.WriteHeader(http.StatusCreated)
		json.NewEncoder(response).Encode(anwser)

	
	}
}


// Function to login
func Security_login(response http.ResponseWriter, request *http.Request) {
	// Set the response header
	response.Header().Set("Content-Type", "application/json")

	// Get the parameters from the request
	var user dto.LoginDto
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		anwser := map[string]string{
			"status":  "error",
			"message": "Invalid request",
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(anwser)
		return
	}

	// Validate if exist record with the same email
	userDatabase := models.User{}
	if database.Database.Where("email = ?", user.Email).Find(&userDatabase).RowsAffected > 0 {
		
		// Compare the password with the hash
		//use bcrypt to compare the password
		passwordBytes := []byte(user.Password)
		passwordBd := []byte(userDatabase.Password)

		err := bcrypt.CompareHashAndPassword(passwordBd, passwordBytes)
		if err != nil {
			anwser := map[string]string{
				"status":  "error",
				"message": "Invalid password",
			}
			response.WriteHeader(http.StatusNotFound)
			json.NewEncoder(response).Encode(anwser)
			return
		} else {
			jwtKey, errJwt := jwt.GenerateJWT(userDatabase)
			if errJwt != nil {
				anwser := map[string]string{
					"status":  "error",
					"message": "Error generating the token" + errJwt.Error(),
				}
				response.WriteHeader(http.StatusNotFound)
				json.NewEncoder(response).Encode(anwser)
				return
			} else {
				anwser := dto.LoginResponseDto{
					Name:  userDatabase.Name,
					Token: jwtKey,
				}
				response.WriteHeader(http.StatusOK)
				json.NewEncoder(response).Encode(anwser)
			}
		}
	} else {
		anwser := map[string]string{
			"status":  "error",
			"message": "User not found",
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(anwser)
		return
	}
}


//Function to confirm the token
func Security_protected(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content-Type", "application/json")
	anwser := map[string]string{
		"status":  "Ok",
		"message": "Route protected",
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(anwser)

}
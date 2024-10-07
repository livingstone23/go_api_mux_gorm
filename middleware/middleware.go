package middleware

import (
	"encoding/json"
	"fmt"
	"go_api_mux_gorm/database"
	"go_api_mux_gorm/models"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// Function to validate the token
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		//Get the existenc os .env file and the variables
		errorVar := godotenv.Load()
		if errorVar != nil {
			panic("Error loading .env file")
		}

		myKey := []byte(os.Getenv("SECRET_JWT"))

		// Set the response header
		response.Header().Set("Content-Type", "application/json")
		header := request.Header.Get("Authorization")
		if len(header) == 0 {
			anwser := map[string]string{
				"status":  "error",
				"message": "Token is required",
			}
			response.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(response).Encode(anwser)
			return
		}

		//Validate if the token is empty
		splitBearer := strings.Split(header, " ")
		if len(splitBearer) != 2 {
			anwser := map[string]string{
				"status":  "error",
				"message": "No Authorization token",
			}
			response.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(response).Encode(anwser)
			return
		}

		//Remove the space from the token
		splitToken := strings.Split(splitBearer[1], ".")
		if len(splitToken) != 3 {
			anwser := map[string]string{
				"status":  "error",
				"message": "No authorization token",
			}
			response.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(response).Encode(anwser)
			return
		}

		//Decode the token
		tk := strings.TrimSpace(splitBearer[1])
		token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return myKey, nil
		})

		if err != nil {
			anwser := map[string]string{
				"status":  "error",
				"message": "No authorization token",
			}
			response.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(response).Encode(anwser)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user := models.User{}
			if err := database.Database.Where("email = ?", claims["email"]).First(&user).Error; err != nil {
				anwser := map[string]string{
					"status":  "error",
					"message": "User not found",
				}
				response.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(response).Encode(anwser)
				return
			} else {
				next.ServeHTTP(response, request)
			}
		} else {
			anwser := map[string]string{
				"status":  "error",
				"message": "No authorization token",
			}
			response.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(response).Encode(anwser)
			return
		}

	})
}

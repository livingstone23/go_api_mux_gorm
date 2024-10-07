package jwt

import (
	"go_api_mux_gorm/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// func to generate the token
func GenerateJWT(user models.User) (string, error) {
	errorVA := godotenv.Load()
	if errorVA != nil {
		panic("Error loading .env file")
	}

	myKey := []byte(os.Getenv("SECRET_JWT"))

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":        user.Email,
		"name":         user.Name,
		"created_from": "API",
		"id":           user.Id,
		"exp":          time.Now().Add(time.Hour * 24).Unix(), //add 24 hours to the token
	})

	tokenString, err := token.SignedString(myKey)
	return tokenString, err
}

package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwtSecret []byte

func init() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        panic("Failed to load .env file")
    }

    // Get JWT secret from environment variable
    jwtSecret = []byte(os.Getenv("JWT_SECRET"))
    if len(jwtSecret) == 0 {
        panic("JWT_SECRET is not set in .env file")
    }
}

// GetJWTSecret returns the JWT secret key
func GetJWTSecret() []byte {
    return jwtSecret
}

func GenerateToken(userID uint) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

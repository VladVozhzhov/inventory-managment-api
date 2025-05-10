package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

var DB *gorm.DB
var JWTSecret string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}
	JWTSecret = os.Getenv("JWT_SECRET")
	log.Printf("DB_USER: %s", os.Getenv("DB_USER"))
	log.Printf("DB_PASSWORD: %s", os.Getenv("DB_PASSWORD"))
}

func ConnectDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	log.Println("Connected to database")

	return db
}

func GenerateJWT(userID string, role string) (string, error) {
	jwtKey := []byte(JWTSecret)
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := struct {
		Role string `json:"role"`
		jwt.RegisteredClaims
	}{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

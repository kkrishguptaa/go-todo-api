package util

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	loadEnv()

	env := os.Environ()

	for _, e := range env {
		fmt.Println("loaded env variable: " + e)
	}

	if getEnv("DATABASE_URL") == "" || getEnv("DATABASE_NAME") == "" {
		log.Panic("env variables DATABASE_URL and DATABASE_NAME are required")
	}
}

func getEnv(key string) string {
	return os.Getenv(key)
}

func isProd() bool {
	return getEnv("PRODUCTION") == "true" || getEnv("PRODUCTION") == "1" || getEnv("RENDER") == "true"
}

func loadEnv() {
	if isProd() {
		fmt.Println("mode: production")
		return
	}

	fmt.Println("mode: development")
	fmt.Println("loading .env file")

	err := godotenv.Load(".env")

	if err != nil {
		log.Panic("error loading .env file")
	}

	fmt.Println("loaded .env file")
}

package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	load()

	env := os.Environ()

	for _, e := range env {
		fmt.Println("loaded env variable: " + e)
	}

	if Get("DATABASE_URL") == "" || Get("DATABASE_NAME") == "" {
		log.Panic("env variables DATABASE_URL and DATABASE_NAME are required")
	}
}

func Get(key string) string {
	return os.Getenv(key)
}

func IsProd() bool {
	return Get("PRODUCTION") == "true" || Get("PRODUCTION") == "1" || Get("RENDER") == "true"
}

func load() {
	if IsProd() {
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

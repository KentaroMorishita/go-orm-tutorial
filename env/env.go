package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(fmt.Sprintf("env/%s.env", Get("GO_ENV", "develop")))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Get env variable by key
func Get(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

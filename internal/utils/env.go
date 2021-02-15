package utils

import (
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

// MustGet will return the env or panic if it is not present
func MustGet(k string) string {
	v := os.Getenv(k)

	if v == "" {
		log.Panicf("ENV is missing key [%s]", k)
	}

	return v
}

// MustGetBool will return the env variable as a bool or panic if not present
func MustGetBool(k string) bool {
	v := os.Getenv(k)

	if v == "" {
		log.Panicf("ENV is missing key [%s]", k)
	}

	b, err := strconv.ParseBool(v)
	if err != nil {
		log.Panicf("ENV err: [%s] %s", k, err.Error())
	}

	return b
}

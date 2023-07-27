package util

import (
	"log"
	"os"
)

// Get environmental variable or panic
func GetEnv(key string) string {
	val, valid := os.LookupEnv(key)
	if !valid || val == "" {
		log.Fatalln(key + " env cannot be found!")
	}

	return val
}

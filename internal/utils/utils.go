package utils

import (
	"log"
	"os"
)

// GetHomeDirectory Define users home directory
func GetHomeDirectory() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return home
}

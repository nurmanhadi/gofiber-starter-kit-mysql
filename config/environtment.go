package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func NewEnvirontment() error {
	envFiles := []string{".env.development", ".env.production"}

	for _, file := range envFiles {
		if _, err := os.Stat(file); err == nil {
			if err := godotenv.Load(file); err != nil {
				log.Fatalf("environtment error: %s", err.Error())
			} else {
				log.Printf("environtment load: %s", file)
				return nil
			}
		}
	}
	log.Fatal("no such or directory for .env.development and .env.production")
	return nil
}

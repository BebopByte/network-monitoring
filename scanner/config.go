package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Subnet string
}

func LoadConfig() Config {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		Subnet: os.Getenv("SUBNET"),
	}
}

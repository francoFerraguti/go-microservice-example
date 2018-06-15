package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration interface {
	Load()
	LoadedCorrectly() bool

	Port() string
}

type configuration struct {
	port string
}

func (c *configuration) Load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("warning: enviromental file not found")
	}

	c.port = os.Getenv("PORT")

	if !c.LoadedCorrectly() {
		log.Fatal("couldn't find configuration")
	}
}

func (c *configuration) LoadedCorrectly() bool {
	return c.port != ""
}

func (c *configuration) Port() string {
	return c.port
}

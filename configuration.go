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
	Database() databaseConfiguration
}

type configuration struct {
	port                  string
	databaseConfiguration databaseConfiguration
}

type databaseConfiguration struct {
	name     string
	host     string
	port     string
	username string
	password string
}

func (c *configuration) Load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("warning: enviromental file not found")
	}

	c.port = os.Getenv("PORT")
	c.databaseConfiguration.name = os.Getenv("DB_NAME")
	c.databaseConfiguration.host = os.Getenv("DB_HOST")
	c.databaseConfiguration.port = os.Getenv("DB_PORT")
	c.databaseConfiguration.username = os.Getenv("DB_USERNAME")
	c.databaseConfiguration.password = os.Getenv("DB_PASSWORD")

	if !c.LoadedCorrectly() {
		log.Fatal("couldn't find configuration")
	}
}

func (c *configuration) LoadedCorrectly() bool {
	return c.port != "" && c.databaseConfiguration.name != "" && c.databaseConfiguration.host != "" &&
		c.databaseConfiguration.port != "" && c.databaseConfiguration.username != "" && c.databaseConfiguration.password != ""
}

func (c *configuration) Port() string {
	return c.port
}

func (c *configuration) Database() databaseConfiguration {
	return c.databaseConfiguration
}

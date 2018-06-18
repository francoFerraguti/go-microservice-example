package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var myConfiguration *configuration

type Configuration interface {
	Load()
	LoadedCorrectly() bool

	Port() string
	ProjectName() string
	Database() *databaseConfiguration
}

type configuration struct {
	port                  string
	projectName           string
	databaseConfiguration *databaseConfiguration
}

func (c *configuration) Load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("warning: enviromental file not found")
	}

	c.port = os.Getenv("PORT")
	c.projectName = os.Getenv("PROJECT_NAME")
	c.databaseConfiguration.name = os.Getenv("DB_NAME")
	c.databaseConfiguration.host = os.Getenv("DB_HOST")
	c.databaseConfiguration.port = os.Getenv("DB_PORT")
	c.databaseConfiguration.username = os.Getenv("DB_USERNAME")
	c.databaseConfiguration.password = os.Getenv("DB_PASSWORD")

	if os.Getenv("DELETE_DATABASE") == "true" {
		c.databaseConfiguration.delete = true
	} else {
		c.databaseConfiguration.delete = false
	}

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

func (c *configuration) ProjectName() string {
	return c.projectName
}

func (c *configuration) Database() *databaseConfiguration {
	return c.databaseConfiguration
}

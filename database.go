package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database interface {
	Get() *sql.DB
	Connect()
}

type database struct {
	db *sql.DB
}

func (d *database) Get() *sql.DB {
	return d.db
}

func (d *database) Connect(c *configuration) {
	var err error
	name := c.Database().name
	host := c.Database().host
	port := c.Database().port
	username := c.Database().username
	password := c.Database().password

	d.db, err = sql.Open("mysql", username+":"+password+"@"+"tcp("+host+":"+port+")"+"/"+name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
	}
}

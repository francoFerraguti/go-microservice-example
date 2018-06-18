package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var myDatabase *database

type Database interface {
	Get() *sql.DB
	Connect()
	Setup()
	Create()
	Delete()
	ExecuteSQLFile(path string)
	NumberOfTables() int
}

type database struct {
	db *sql.DB
}

func (d *database) Get() *sql.DB {
	return d.db
}

func (d *database) Connect() {
	var err error
	name := myConfiguration.Database().Name()
	host := myConfiguration.Database().Host()
	port := myConfiguration.Database().Port()
	username := myConfiguration.Database().Username()
	password := myConfiguration.Database().Password()

	d.db, err = sql.Open("mysql", username+":"+password+"@"+"tcp("+host+":"+port+")"+"/"+name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = d.db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (d *database) Setup() {
	if myConfiguration.Database().Delete() {
		d.Delete()
	}

	d.Create()
}

func (d *database) ExecuteSQLFile(path string) {
	sqlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	requests := strings.Split(string(sqlFile), ";")

	for _, request := range requests {
		if request == "" {
			continue
		}

		_, err := d.db.Exec(request)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func (d *database) Delete() {
	d.ExecuteSQLFile("delete.sql")
}

func (d *database) Create() {
	d.ExecuteSQLFile("create.sql")
}

func (d *database) NumberOfTables() int {
	var numberOfTables int
	d.db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = '" + myConfiguration.Database().Name() + "';").Scan(&numberOfTables)
	return numberOfTables
}

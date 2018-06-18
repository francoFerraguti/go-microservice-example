package main

type DatabaseConfiguration interface {
	Name() string
	Host() string
	Port() string
	Username() string
	Password() string
	Delete() bool
}

type databaseConfiguration struct {
	name     string
	host     string
	port     string
	username string
	password string
	delete   bool
}

func (db *databaseConfiguration) Name() string {
	return db.name
}

func (db *databaseConfiguration) Host() string {
	return db.host
}

func (db *databaseConfiguration) Port() string {
	return db.port
}

func (db *databaseConfiguration) Username() string {
	return db.username
}

func (db *databaseConfiguration) Password() string {
	return db.password
}

func (db *databaseConfiguration) Delete() bool {
	return db.delete
}

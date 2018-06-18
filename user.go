package main

import "time"

type User interface {
	// Data
	ID() int
	Username() string
	Email() string
	Password() string
	Enabled() bool
	DateCreated() time.Time

	// Model
	Create() (*user, error)
	TableExists() bool
}

type user struct {
	id          int
	username    string
	email       string
	password    string
	enabled     bool
	dateCreated time.Time
}

func (u *user) ID() int {
	return u.id
}

func (u *user) Username() string {
	return u.username
}

func (u *user) Email() string {
	return u.email
}

func (u *user) Password() string {
	return u.password
}

func (u *user) Enabled() bool {
	return u.enabled
}

func (u *user) DateCreated() time.Time {
	return u.dateCreated
}

package main

import (
	"crypto/sha1"
	"encoding/base64"
)

var myHelper *helper

type Helper interface {
	HashPassword(password string, salt string) string
}

type helper struct {
}

func (h *helper) HashPassword(password string, salt string) string {
	hasher := sha1.New()
	hasher.Write([]byte(salt + password))
	hashedPassword := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return hashedPassword
}

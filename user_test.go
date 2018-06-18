package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableExists(t *testing.T) {
	SetupProject()

	user := &user{}
	assert.True(t, user.TableExists())
}

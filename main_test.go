package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	SetupProject()
	myRouter.Init()

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/v1/ping", nil)
	assert.NoError(t, err)
	myRouter.Get().ServeHTTP(recorder, request)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "pong", recorder.Body.String())
}

func TestConfiguration(t *testing.T) {
	SetupProject()
	myConfiguration.Load()

	assert.Equal(t, os.Getenv("PORT"), myConfiguration.Port())
	assert.Equal(t, os.Getenv("PROJECT_NAME"), myConfiguration.ProjectName())
	assert.Equal(t, os.Getenv("DB_NAME"), myConfiguration.Database().Name())
	assert.Equal(t, os.Getenv("DB_HOST"), myConfiguration.Database().Host())
	assert.Equal(t, os.Getenv("DB_PORT"), myConfiguration.Database().Port())
	assert.Equal(t, os.Getenv("DB_USERNAME"), myConfiguration.Database().Username())
	assert.Equal(t, os.Getenv("DB_PASSWORD"), myConfiguration.Database().Password())
}

func TestDatabase(t *testing.T) {
	SetupProject()
	myConfiguration.Load()

	myDatabase.Connect()
	myDatabase.Setup()
	assert.Equal(t, 1, myDatabase.NumberOfTables())
}

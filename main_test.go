package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := router{}
	router.Init()

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/v1/ping", nil)
	assert.NoError(t, err)
	router.Get().ServeHTTP(recorder, request)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "pong", recorder.Body.String())
}

func TestConfiguration(t *testing.T) {
	configuration := configuration{}
	configuration.Load()

	assert.Equal(t, os.Getenv("PORT"), configuration.Port())
}

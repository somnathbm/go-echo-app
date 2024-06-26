package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// test health check function
func TestHealthy(t *testing.T) {
	// mock `/healthy` response
	mockedHealthyResponseJSON := `{"status": "OK"}`

	// test setup

	// request comes in
	req := httptest.NewRequest(http.MethodGet, "/healthy", nil)

	// set the request header
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// simple assertion
	assert := assert.New(t)
	assert.Equal(http.StatusOK, rec.Code)
	assert.Equal(mockedHealthyResponseJSON, `{"status": "OK"}`)
}

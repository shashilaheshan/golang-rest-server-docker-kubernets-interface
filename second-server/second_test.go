package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	// Create a request
	req, err := http.NewRequest(http.MethodGet, "/books", nil)
	assert.NoError(t, err)

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler function
	getBooks(rr, req)

	// Validate the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Validate response body
	expectedBooks := []Books{
		{Name: "Shashila Heshan"},
	}

	var actualBooks []Books
	err = json.Unmarshal(rr.Body.Bytes(), &actualBooks)
	assert.NoError(t, err)

	assert.Equal(t, expectedBooks, actualBooks)
}
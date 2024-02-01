package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type FailingMarshaler struct{}

func (fm *FailingMarshaler) MarshalJSON() ([]byte, error) {
	return nil, errors.New("forced marshal error")
}

func TestResponseWithJSON(t *testing.T) {
	w := httptest.NewRecorder()
	payload := map[string]string{
		"message": "Test",
	}

	ResponseWithJSON(w, 201, payload)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	// Check the Content-Type header
	expectedContentType := "Application/json"
	actualContentType := w.Header().Get("Content-Type")
	if actualContentType != expectedContentType {
		t.Errorf("Expected Content-Type %s, got %s", expectedContentType, actualContentType)
	}

	// Decode the response body and check its content
	var responseBody map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	expectedMessage := "Test"
	actualMessage, exists := responseBody["message"]
	if !exists || actualMessage != expectedMessage {
		t.Errorf("Expected message %s, got %s", expectedMessage, actualMessage)
	}

}

func TestResponseWithJSON_MarshalError(t *testing.T) {
	// Create a mock response writer
	w := httptest.NewRecorder()

	// Use a failing marshaler
	payload := &FailingMarshaler{}

	// Call the function directly
	ResponseWithJSON(w, http.StatusOK, payload)

	// Check the response status code
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
	}

	// Check that the response body is empty (or contains an error message)
	if w.Body.Len() > 0 {
		t.Errorf("Expected empty response body, got %s", w.Body.String())
	}
}

func TestResponseWithError(t *testing.T) {
	w := httptest.NewRecorder()
	ResponseWithError(w, 501, "This test should fail")

	if w.Code < 500 {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)

	}
}

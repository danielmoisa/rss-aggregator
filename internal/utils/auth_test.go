package utils

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	// Test case 1: Valid Authorization header
	header := http.Header{"Authorization": []string{"ApiKey myapikey"}}

	apiKey, err := GetApiKey(header)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedApiKey := "myapikey"
	if apiKey != expectedApiKey {
		t.Errorf("Expected API key '%s', got '%s'", expectedApiKey, apiKey)
	}

	// Test case 2: Missing header
	header = http.Header{}

	apiKey, err = GetApiKey(header)
	if err == nil || apiKey != "" {
		t.Errorf("Expected error, got none. Expected empty API key, got '%s'", apiKey)
	}

	// Test case 3: Malformed Authorization header
	header = http.Header{"Authorization": []string{"malformed"}}
	apiKey, err = GetApiKey(header)
	if err == nil || apiKey != "" {
		t.Errorf("Expected error, got none. Expected empty API key, got '%s'", apiKey)
	}

	// Test case 4: Invalid auth header format
	header = http.Header{"Authorization": []string{"Bearer invalid"}}
	apiKey, err = GetApiKey(header)
	if err == nil || apiKey != "" {
		t.Errorf("Expected error, got none. Expected empty API key, got '%s'", apiKey)
	}

}

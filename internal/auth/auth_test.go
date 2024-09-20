package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name            string
		inputAuthHeader string
		expected        string
		expectedErr     bool
	}{
		{
			name:            "Test no Authorization header",
			inputAuthHeader: "",
			expected:        "",
			expectedErr:     true,
		},
		{
			name:            "Test valid api key",
			inputAuthHeader: "ApiKey testapikey",
			expected:        "testapikey",
			expectedErr:     false,
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			headers := http.Header{}
			if tc.inputAuthHeader != "" {
				headers.Add("Authorization", tc.inputAuthHeader)
			}
			actual, err := GetAPIKey(headers)
			if err == nil && tc.expectedErr {
				t.Errorf("Test %v - '%s' FAIL: expected an error none were returned", i, tc.name)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - '%s' FAIL: Expected ApiKey: %s; Actual: %s", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}

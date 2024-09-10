package search

import (
	"io"
	"strings"
	"testing"
)

func TestParseHTML(t *testing.T) {
	tests := []struct {
		name            string
		htmlInput       string
		expectedTitle   string
		expectedContent string
		expectError     bool
	}{
		{
			name: "Valid HTML with Title and Content",
			htmlInput: `
				<!DOCTYPE html>
				<html>
				<head><title>Test Page</title></head>
				<body>
					<h1 class="inner-name">Test Article Title</h1>
					<div class="inner-content">This is the content of the test article.</div>
				</body>
				</html>
			`,
			expectedTitle:   "Test Article Title",
			expectedContent: "This is the content of the test article.",
			expectError:     false,
		},
		{
			name: "HTML with Missing Content",
			htmlInput: `
				<!DOCTYPE html>
				<html>
				<head><title>Test Page</title></head>
				<body>
					<h1 class="inner-name">Test Article Title</h1>
				</body>
				</html>
			`,
			expectedTitle:   "Test Article Title",
			expectedContent: "",
			expectError:     false,
		},
		{
			name: "Empty HTML Document",
			htmlInput: `
				<!DOCTYPE html>
				<html>
				<head><title>Test Page</title></head>
				<body>
				</body>
				</html>
			`,
			expectedTitle:   "",
			expectedContent: "",
			expectError:     false,
		},
		{
			name:            "Invalid HTML Input",
			htmlInput:       `<html><head><title>Missing body tag`,
			expectedTitle:   "",
			expectedContent: "",
			expectError:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reader io.Reader = strings.NewReader(tt.htmlInput)

			title, content, err := ParseHTML(reader)

			if tt.expectError && err == nil {
				t.Errorf("Expected error but got none")
			} else if !tt.expectError && err != nil {
				t.Errorf("Did not expect error but got one: %v", err)
			}

			if title != tt.expectedTitle {
				t.Errorf("Expected title '%s' but got '%s'", tt.expectedTitle, title)
			}

			if content != tt.expectedContent {
				t.Errorf("Expected content '%s' but got '%s'", tt.expectedContent, content)
			}
		})
	}
}

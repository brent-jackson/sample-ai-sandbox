package main

import (
	"testing"
)

func TestPerformCalculation(t *testing.T) {
	tests := []struct {
		expression string
		expected   string
		shouldErr  bool
	}{
		{"2 + 3", "5", false},
		{"10 - 4", "6", false},
		{"3 * 7", "21", false},
		{"15 / 3", "5", false},
		{"sqrt(16)", "4", false},
		{"sin(pi/2)", "1", false},
		{"invalid", "", true},
	}

	for _, test := range tests {
		result, err := performCalculation(test.expression)
		
		if test.shouldErr {
			if err == nil {
				t.Errorf("Expected error for expression %s, but got none", test.expression)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for expression %s: %v", test.expression, err)
			}
			if result != test.expected {
				t.Errorf("For expression %s, expected %s, got %s", test.expression, test.expected, result)
			}
		}
	}
}

func TestTransformText(t *testing.T) {
	tests := []struct {
		text      string
		operation string
		expected  string
		shouldErr bool
	}{
		{"hello", "uppercase", "HELLO", false},
		{"HELLO", "lowercase", "hello", false},
		{"hello", "reverse", "olleh", false},
		{"hello world", "word_count", "Word count: 2", false},
		{"hello", "char_count", "Character count: 5", false},
		{"hello", "invalid", "", true},
	}

	for _, test := range tests {
		result, err := transformText(test.text, test.operation)
		
		if test.shouldErr {
			if err == nil {
				t.Errorf("Expected error for operation %s, but got none", test.operation)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for operation %s: %v", test.operation, err)
			}
			if result != test.expected {
				t.Errorf("For operation %s on text %s, expected %s, got %s", test.operation, test.text, test.expected, result)
			}
		}
	}
}

func TestGetCurrentTime(t *testing.T) {
	tests := []struct {
		format       string
		customFormat string
		shouldErr    bool
	}{
		{"iso8601", "", false},
		{"unix", "", false},
		{"human_readable", "", false},
		{"custom", "2006-01-02", false},
		{"custom", "", true}, // custom without custom_format should error
		{"invalid", "", true},
	}

	for _, test := range tests {
		result, err := getCurrentTime(test.format, test.customFormat)
		
		if test.shouldErr {
			if err == nil {
				t.Errorf("Expected error for format %s, but got none", test.format)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for format %s: %v", test.format, err)
			}
			if result == "" {
				t.Errorf("Expected non-empty result for format %s", test.format)
			}
		}
	}
}
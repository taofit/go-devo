package main

import "testing"

func TestFormatWithCommas(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, "0"},
		{5, "5"},
		{12, "12"},
		{123, "123"},
		{1000, "1,000"},
		{12345, "12,345"},
		{1234567, "1,234,567"},
		{1234567890, "1,234,567,890"},
	}

	for _, tt := range tests {
		result := FormatWithCommas(tt.input)
		if result != tt.expected {
			t.Errorf("FormatWithCommas(%d) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

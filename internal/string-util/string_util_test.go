package stringutil_test

import (
	stringutil "github.com/GokselKUCUKSAHIN/jsonx/internal/string-util"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_Byte(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []byte
	}{
		{"Test 1", "hello", []byte{104, 101, 108, 108, 111}},
		{"Test 2", "world", []byte{119, 111, 114, 108, 100}},
		{"Test 3", "", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stringutil.Byte(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Test: %s, Got: %v, Expected: %v", tt.name, result, tt.expected)
			}
		})
	}
}

func Test_String(t *testing.T) {
	tests := []struct {
		name  string
		bytes []byte
	}{
		{
			name:  "Empty slice",
			bytes: []byte{},
		},
		{
			name:  "Single character slice",
			bytes: []byte{'a'},
		},
		{
			name:  "Multiple characters slice",
			bytes: []byte{'h', 'e', 'l', 'l', 'o'},
		},
		{
			name:  "Slice with spaces",
			bytes: []byte{' ', ' ', ' '},
		},
		{
			name:  "Slice with numbers",
			bytes: []byte{'1', '2', '3'},
		},
		{
			name:  "Slice with special characters",
			bytes: []byte{'!', '@', '#'},
		},
		{
			name:  "Slice with mixed characters",
			bytes: []byte{'H', 'e', 'l', 'l', 'o', ' ', '1', '2', '3', '!'},
		},
		{
			name:  "Special Characters",
			bytes: []byte("!@#$%^&*()_+-=[]{}|;':\",./<>?"),
		},
		{
			name:  "Unicode Characters",
			bytes: []byte("こんにちは世界"),
		},
		{
			name:  "Modify Original Slice After Conversion",
			bytes: []byte("Mutable"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := string(tt.bytes)
			result := stringutil.String(tt.bytes)

			// Check if the result from the String function matches the built-in conversion
			assert.Equal(t, expected, result, "Result from String function should match built-in conversion")
		})
	}
}

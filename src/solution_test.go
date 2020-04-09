package main

import "testing"

func TestReverse(t *testing.T) {
	tests := map[string]string{
		"A4B5C2":   "AAAABBBBBCC",
		"C2F1E5":   "CCFEEEEE",
		"T4S2V2":   "TTTTSSVV",
		"A1B2C3D4": "ABBCCCDDDD",
	}

	for key, element := range tests {
		output := reverse(key)

		if element != output {
			t.Errorf("Reversed output for '%s' was incorrect, expected '%s', got '%s'", key, element, output)
		}
	}
}

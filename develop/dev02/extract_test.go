package main

import (
	"testing"
	"strings"
)

func TestExtract(t *testing.T) {
	tests := map[string]string{
		"a4bc2d5e": "aaaabccddddde",
		"abcd":     "abcd",
		"45":       "Incorrect string",
		"":         "",
	}

	for key, value := range tests {
		res := extract(key)

		if strings.Compare(res, value) != 0 {
			t.Errorf("Error incorrect string; current %s, expected %s", res, value)
			continue
		}
	}
}

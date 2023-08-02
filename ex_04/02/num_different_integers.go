package main

import (
	"strconv"
	"strings"
	"unicode"
)

// numDifferentIntegers returns a different integers in a word
func numDifferentIntegers(word string) int {
	numMap := make(map[int]bool)
	numStr := strings.Builder{}

	for _, char := range word {
		if unicode.IsDigit(char) {
			numStr.WriteRune(char)
		} else if numStr.Len() > 0 {
			if key, err := strconv.Atoi(numStr.String()); err == nil {
				numMap[key] = true
				numStr.Reset()
			}
		}
	}

	// Check if there is an integer at the end of the word
	if numStr.Len() > 0 {
		if key, err := strconv.Atoi(numStr.String()); err == nil {
			numMap[key] = true
		}
	}

	return len(numMap)
}

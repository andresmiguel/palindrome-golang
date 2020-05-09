package main

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	var letters []rune

	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	size := len(letters) / 2
	for i := 0; i < size; i++ {
		if letters[i] != letters[len(letters)-i-1] {
			return false
		}
	}

	return true
}

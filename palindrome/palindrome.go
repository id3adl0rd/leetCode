package palindrome

import (
	"strings"
	"unicode"
)

func lowering(r rune) rune {
	if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
		return -1
	}

	return unicode.ToLower(r)
}

func isPalindrome(s string) bool {
	s = strings.Map(lowering, s)

	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			return false
		}

		l = l + 1
		r = r - 1
	}

	return true
}

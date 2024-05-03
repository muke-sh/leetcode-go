package longestpalindrome

import "fmt"

func LongestPalindrome(s string) string {

	strLength := len(s)

	if strLength == 1 {
		return s
	}

	palindromeSubstring := make(map[string]bool)

	longestPalidromeSoFar := fmt.Sprintf("%c", s[0])

	for i := 0; i < strLength; i++ {
		for j := i + 1; j < strLength; j++ {
			substring := s[i : j+1]

			if len(substring) > len(longestPalidromeSoFar) &&
				!isSubstringAlreadyFound(palindromeSubstring, substring) &&
				IsPalindrome(substring) {
				longestPalidromeSoFar = substring
				palindromeSubstring[substring] = true
			}
		}
	}

	return longestPalidromeSoFar
}

func isSubstringAlreadyFound(palindromeSubstring map[string]bool, str string) bool {

	if _, ok := palindromeSubstring[str]; ok {
		return true
	}

	return false
}

func IsPalindrome(s string) bool {
	temp := Reverse(s)

	return temp == s
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

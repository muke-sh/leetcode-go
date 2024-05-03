package longestsubstring

import "fmt"

func LengthOfLongestSubstring(inputString string) int {

	strLength := len(inputString)

	longestLengthSoFar := 1

	for i := 0; i < strLength; i++ {
		charEncounteredSoFar := make(map[string]bool)

		currentChar := fmt.Sprintf("%c", inputString[i])
		charEncounteredSoFar[currentChar] = true
		currentLongest := 1

		for j := i + 1; j < strLength; j++ {
			nextChar := fmt.Sprintf("%c", inputString[j])

			if isCharacterSeen(charEncounteredSoFar, nextChar) {
				currentLongest = 1
				break
			} else {
				currentLongest += 1
				charEncounteredSoFar[nextChar] = true
			}

			if currentLongest > longestLengthSoFar {
				longestLengthSoFar = currentLongest
			}
		}
	}
	return longestLengthSoFar
}

func isCharacterSeen(charEncounteredSoFar map[string]bool, char string) bool {

	if _, ok := charEncounteredSoFar[char]; ok {
		return true
	} else {
		return false
	}
}

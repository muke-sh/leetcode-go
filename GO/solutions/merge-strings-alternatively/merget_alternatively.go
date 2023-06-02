package mergestringsalternatively

func mergeAlternately(word1 string, word2 string) string {
	word1Len := len(word1)
	word2Len := len(word2)

	startWord1 := 0
	startWord2 := 0

	endWord1 := word1Len - 1
	endWord2 := word2Len - 1

	var newString string

	for startWord1 <= endWord1 && startWord2 <= endWord2 {
		newString += string(word1[startWord1])
		newString += string(word2[startWord2])
		startWord1 += 1
		startWord2 += 1
	}

	if startWord1 < word1Len {
		newString += word1[startWord1:]
	}

	if startWord2 < word2Len {
		newString += word2[startWord2:]
	}

	return newString
}

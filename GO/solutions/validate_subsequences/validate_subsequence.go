package validate_subsequences

// Time complexity: O(n)
// Space complexity: O(1)

// func validateSubsequence(arr, subsequence []int) bool {

// 	arrIdx, seqIdx := 0, 0

// 	for arrIdx < len(arr) && seqIdx < len(subsequence) {
// 		if arr[arrIdx] == subsequence[seqIdx] {
// 			seqIdx += 1
// 		}
// 		arrIdx += 1
// 	}

// 	return seqIdx == len(subsequence)
// }

func validateSubsequence(arr, subsequence []int) bool {

	seqIdx := 0

	for _, elem := range arr {
		if seqIdx == len(subsequence) {
			return true
		}

		if elem == subsequence[seqIdx] {
			seqIdx += 1
		}
	}

	return seqIdx == len(subsequence)
}

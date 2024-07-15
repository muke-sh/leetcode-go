package consecutive_odd

func threeConsecutiveOdds(arr []int) bool {
	windowSize := 3
	sz := len(arr)
	for l := 0; l < sz; l++ {
		bound := l + windowSize
		if bound > sz {
			return false
		}
		consecArr := arr[l:bound]
		if isConsec(consecArr) {
			return true
		}
	}

	return false
}

func isConsec(consecArr []int) bool {
	return (consecArr[0]%2 != 0) && (consecArr[1]%2 != 0) && (consecArr[2]%2 != 0)
}

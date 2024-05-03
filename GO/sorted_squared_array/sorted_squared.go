package sorted_squaredarray

import "math"

func SortedSquaredArray(arr []int) []int {
	n := len(arr)
	start := 0
	end := n - 1
	squared_arr := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		if math.Abs(float64(arr[start])) > math.Abs(float64(arr[end])) {
			squared_num := arr[start] * arr[start]
			squared_arr[i] = squared_num
			start++
		} else {
			squared_num := arr[end] * arr[end]
			squared_arr[i] = squared_num
			end--
		}
	}
	return squared_arr
}

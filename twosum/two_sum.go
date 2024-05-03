package twosum

// TwoSum https://leetcode.com/problems/two-sum/
func TwoSum(nums []int, target int) []int {

	sum_found := false

	var indices []int = make([]int, 2)

	length := len(nums)

	if length == 2 {
		return []int{0, 1}
	}

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				indices[0] = i
				indices[1] = j

				sum_found = true
				break
			}
		}

		if sum_found {
			break
		}
	}

	return indices
}

func TwoSumHashMap(nums []int, target int) []int {

	hm := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		currenVal := nums[i]

		if val, ok := hm[target-currenVal]; ok {
			return []int{val, i}
		}

		hm[nums[i]] = i
	}

	return []int{-1, -1}
}

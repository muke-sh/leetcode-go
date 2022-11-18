package twosum

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

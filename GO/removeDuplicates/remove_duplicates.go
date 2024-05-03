package removeduplicates

//removeDuplicates: https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	i := 1
	nextNonDuplicate := 1

	for i < len(nums) {
		if nums[nextNonDuplicate-1] != nums[i] {
			nums[nextNonDuplicate] = nums[i]
			nextNonDuplicate++
		}
		i++
	}

	return nextNonDuplicate
}

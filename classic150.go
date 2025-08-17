package main

func removeDuplicates(nums []int) int {
	idx := -1
	for i := range nums {
		if idx < 0 || nums[idx] != nums[i] {
			idx++
			nums[idx] = nums[i]
		}
	}
	return idx + 1
}

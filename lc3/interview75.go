package lc3

import "slices"

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stk := []int{}
	for i := range temperatures {
		for len(stk) > 0 && temperatures[stk[len(stk)-1]] < temperatures[i] {
			lastDay := stk[len(stk)-1]
			ans[lastDay] = i - lastDay
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}
	return ans
}

func longestOnes(nums []int, k int) (ans int) {
	l := 0
	count := 0
	for i := range nums {
		if nums[i] == 0 {
			count++
		}
		for count > k {
			if nums[l] == 0 {
				count--
			}
			l++
		}
		ans = max(ans, i-l+1)
	}
	return
}

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-2
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

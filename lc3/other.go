package main

import (
	"fmt"
	"slices"
)

func MaxValueLessN2(n int, nums []int) int {
	var ans int
	var dfs func(idx int, sum int, match bool)
	nl, x := []int{}, n
	for x > 0 {
		nl = append(nl, x%10)
		x /= 10
	}
	slices.Reverse(nl)
	dfs = func(idx int, sum int, match bool) {
		if idx == len(nl) {
			if sum < n {
				ans = max(ans, sum)
			}
			return
		}
		for i := range nums {
			if match {
				if nums[i] == nl[idx] {
					dfs(idx+1, sum*10+nums[i], true)
				} else if nums[i] < nl[idx] {
					dfs(idx+1, sum*10+nums[i], false)
				}
			} else {
				curSum := sum*10 + nums[i]
				if curSum < n {
					dfs(idx+1, curSum, false)
				}
			}
		}
	}
	dfs(0, 0, true)
	return ans
}

func main() {
	nums := []int{1, 0, 9, 8}
	ans := MaxValueLessN2(999, nums)
	fmt.Println(ans)
}

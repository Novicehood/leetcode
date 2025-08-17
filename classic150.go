package main

import (
	"strconv"
	"strings"
)

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

func restoreIpAddresses(s string) (ans []string) {
	n := len(s)
	path := []string{}
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == n && len(path) == 4 {
			ip := strings.Join(path, ".")
			ans = append(ans, ip)
			return
		}

		for i := idx; i < n; i++ {
			cStr := s[idx : i+1]
			num, _ := strconv.Atoi(cStr)
			if num == 0 {
				path = append(path, "0")
				dfs(i + 1)
				path = path[:len(path)-1]
				return
			}

			if len(path) >= 4 {
				return
			}

			if num > 0 && num <= 255 {
				path = append(path, cStr)
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return
}

package main

import (
	"slices"
	"strconv"
)

func atMostNGivenDigitSet(digits []string, n int) int {
	digits = []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	n = 940860624
	nums := []int{}
	for i := range digits {
		num, _ := strconv.Atoi(digits[i])
		nums = append(nums, num)
	}
	slices.Sort(nums)
	res := 0
	var dfs func(sum int)
	dfs = func(sum int) {
		if sum > n {
			return
		}
		res++
		for i := range nums {
			dfs(sum*10 + nums[i])
		}
	}
	dfs(0)
	return res - 1
}

func quickSort(arr []int, low, high int) {
	if low < high {
		// 进行分区操作，返回基准值的最终位置
		pivotIndex := partition2(arr, low, high)
		// 递归排序左半部分
		quickSort(arr, low, pivotIndex-1)
		// 递归排序右半部分
		quickSort(arr, pivotIndex+1, high)
	}
}

// 分区函数 - 经典交换法
func partition2(arr []int, low, high int) int {
	// 选择最右边的元素作为基准
	pivot := arr[high]

	// 初始化左右指针
	left := low
	right := high - 1

	// 主循环：左右指针向中间移动
	for left <= right {
		// 从左向右找第一个大于等于基准的元素
		for left <= right && arr[left] < pivot {
			left++
		}

		// 从右向左找第一个小于等于基准的元素
		for left <= right && arr[right] > pivot {
			right--
		}

		// 如果左右指针没有相遇，交换元素
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	// 将基准元素交换到正确的位置（left指针位置）
	arr[left], arr[high] = arr[high], arr[left]
	return left
}

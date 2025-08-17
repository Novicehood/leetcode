package main

import (
	"container/heap"
	"maps"
	"math"
	"slices"
	"sort"
)

// 只写hot100的题目

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if idx, ok := numsMap[target-nums[i]]; ok {
			return []int{idx, i}
		}
		numsMap[nums[i]] = i
	}
	return nil
}

func groupAnagrams(strs []string) [][]string {
	resMap := map[string][]string{}
	for _, str := range strs {
		strBytes := []byte(str)
		slices.Sort(strBytes)
		k := string(strBytes)
		if resMap[k] == nil {
			resMap[k] = []string{}
		}
		resMap[k] = append(resMap[k], str)
	}
	return slices.Collect(maps.Values(resMap))
}

//TODO少了三数之和和接雨水

func longestConsecutive(nums []int) (ans int) {
	countMap := map[int]int{}
	for _, num := range nums {
		countMap[num]++
	}
	for k, _ := range countMap {
		if countMap[k-1] > 0 {
			continue
		}
		y := k
		for countMap[y] > 0 {
			y++
		}
		ans = max(y-k, ans)
	}
	return
}

func moveZeroes(nums []int) {
	var idx0 int
	for idx, num := range nums {
		if num != 0 {
			nums[idx0], nums[idx] = nums[idx], nums[idx0]
			idx0++
		}
	}
}

func maxArea(height []int) (ans int) {
	left, right := 0, len(height)-1
	for left < right {
		ans = max(ans, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return
}

func lengthOfLongestSubstring(s string) (ans int) {
	sbytes := []byte(s)
	countMap := map[byte]int{}
	left := 0
	for right, b := range sbytes {
		countMap[b]++
		for countMap[b] > 1 {
			countMap[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func findAnagrams(s string, p string) (ans []int) {
	cnp := map[byte]int{}
	for _, ch := range []byte(p) {
		cnp[ch]++
	}
	cns := map[byte]int{}
	left := -1
	for right, ch := range []byte(s) {
		cns[ch]++
		for cns[ch] > cnp[ch] {
			left++
			cns[s[left]]--
		}
		if right-left == len(p) {
			ans = append(ans, left+1)
		}
	}
	return
}

func subarraySum(nums []int, k int) (ans int) {
	sumMap := map[int]int{}
	sum := 0
	sumMap[0]++
	for _, num := range nums {
		sum += num
		ans += sumMap[sum-k]
		sumMap[sum]++
	}
	return
}

func maxSlidingWindow(nums []int, k int) (ans []int) {
	q := []int{}
	for i, num := range nums {
		for len(q) > 0 && num >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i-q[0] > k-1 {
			q = q[1:]
		}
		if i >= k-1 {
			ans = append(ans, nums[q[0]])
		}
	}
	return
}

func minWindow(s string, t string) string {
	chrMap := map[byte]int{}
	chrType := 0
	for i, _ := range t {
		chrMap[t[i]]++
		if chrMap[t[i]] == 1 {
			chrType++
		}
	}
	ansl, ansr, left := 0, len(s)-1, 0
	for i, _ := range s {
		chrMap[s[i]]--
		if chrMap[s[i]] == 0 {
			chrType--
		}
		for chrType == 0 {
			if ansr-ansl > i-left {
				ansr, ansl = i, left
			}
			chrMap[s[left]]++
			if chrMap[s[left]] == 1 {
				chrType++
			}
			left++
		}
	}
	if ansr-ansl == len(s)-1 {
		return ""
	}
	return s[ansl : ansr+1]
}

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+dp[i-1] < nums[i] {
			dp[i] = nums[i]
		} else {
			dp[i] = nums[i] + dp[i-1]
		}
	}
	return slices.Max(dp)
}

func merge(intervals [][]int) (ans [][]int) {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	start, end := intervals[0][0], intervals[0][1]
	for _, interval := range intervals {
		if interval[0] <= end {
			end = max(end, interval[1])
		} else {
			ans = append(ans, []int{start, end})
			start, end = interval[0], interval[1]
		}
	}
	ans = append(ans, []int{start, end})
	return
}

func rotate(nums []int, k int) {
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	prefix[0], suffix[len(nums)-1] = nums[0], nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] * nums[i]
		pidx := len(nums) - 1 - i
		suffix[pidx] = nums[pidx] * suffix[pidx+1]
	}
	for i := 1; i < len(nums)-1; i++ {
		ans[i] = prefix[i-1] * suffix[i+1]
	}
	ans[0], ans[len(nums)-1] = suffix[0], prefix[len(nums)-2]
	return ans
}

func spiralOrder(matrix [][]int) (ans []int) {
	direct := [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
	x, y, i := 0, -1, 0
	dx := 0
	count := len(matrix) * len(matrix[0])
	for i < count {
		nx, ny := x+direct[dx][0], y+direct[dx][1]
		if nx < 0 || nx >= len(matrix) || ny < 0 || ny >= len(matrix[0]) || matrix[nx][ny] == -101 {
			dx++
			dx = dx % 4
		} else {
			i++
			ans = append(ans, matrix[nx][ny])
			matrix[nx][ny] = -101
			x, y = nx, ny
		}
	}
	return
}

func searchMatrix1(matrix [][]int, target int) bool {
	row, col := 0, len(matrix[0])-1
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	tempA, tempB := headA, headB
	for tempA != nil {
		lenA++
		tempA = tempA.Next
	}
	for tempB != nil {
		lenB++
		tempB = tempB.Next
	}
	if lenA == 0 || lenB == 0 {
		return nil
	}
	if lenA > lenB {
		var i int
		for i < lenA-lenB {
			i++
			headA = headA.Next
		}
	}
	if lenA < lenB {
		var i int
		for i < lenB-lenA {
			i++
			headB = headB.Next
		}
	}
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, cur, nxt *ListNode = nil, head, nil
	for cur != nil {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func isPalindromeNode(head *ListNode) bool {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-1-i] {
			return false
		}
	}
	return true
}

func hasCycle(head *ListNode) bool {
	dummy := &ListNode{}
	dummy.Next = head
	fast, slow := dummy, dummy
	for fast != nil && fast.Next != nil && slow != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow && slow != nil {
			return true
		}
	}
	return false
}

var direct = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}

func numIslandsBFS(grid [][]byte, x, y int) {
	if grid[x][y] == '0' || grid[x][y] == '2' {
		return
	}
	grid[x][y] = '2'
	for i := range direct {
		nx, ny := direct[i][0]+x, direct[i][1]+y
		if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) {
			numIslandsBFS(grid, nx, ny)
		}
	}
}

func numIslands(grid [][]byte) (ans int) {
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == '1' {
				numIslandsBFS(grid, x, y)
				ans++
			}
		}
	}
	return
}

func trap(height []int) (ans int) {
	l, r := 0, 0
	lmx, rmx := make([]int, len(height)), make([]int, len(height))
	for i := 0; i < len(height); i++ {
		l = max(l, height[i])
		lmx[i] = l
		r = max(r, height[len(height)-i-1])
		rmx[len(height)-1-i] = r
	}
	for i := range height {
		ans += min(lmx[i], rmx[i]) - height[i]
	}
	return
}

func orangesRotting(grid [][]int) (round int) {
	num := 0
	queue := [][2]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				num++
			} else if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	for len(queue) > 0 && num > 0 {
		round++
		count := len(queue)
		for i := 0; i < count; i++ {
			x, y := queue[0][0], queue[0][1]
			queue = queue[1:]
			for i := range direct {
				nx, ny := x+direct[i][0], y+direct[i][1]
				if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) {
					if grid[nx][ny] == 1 {
						num--
						grid[nx][ny] = 2
						queue = append(queue, [2]int{nx, ny})
					}
				}
			}
		}
	}
	if num > 0 {
		return -1
	}
	return
}

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func permute(nums []int) [][]int {
	ans := [][]int{}
	path := make([]int, len(nums))
	flag := make([]bool, len(nums))
	var dfs func([]int, int)
	dfs = func(nums []int, idx int) {
		if idx == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := range nums {
			if flag[i] {
				continue
			}
			flag[i] = true
			path[idx] = nums[i]
			dfs(nums, idx+1)
			flag[i] = false
		}
	}
	dfs(nums, 0)
	return ans
}

func lowerBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
func searchRange(nums []int, target int) []int {
	le, lg := lowerBound(nums, target), lowerBound(nums, target+1)
	if le >= len(nums) || nums[le] != target {
		return []int{-1, -1}
	}
	return []int{le, lg - 1}
}

func rob(nums []int) int {
	f := make([]int, len(nums)+2)
	for i := range nums {
		f[i+2] = max(f[i]+nums[i], f[i+1])
	}
	return f[len(f)-1]
}

func coinChange(coins []int, amount int) int {
	const VALUE = math.MaxInt/3 - 1
	f := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		f[i] = VALUE
		for j := range coins {
			if i-coins[j] < 0 {
				continue
			}
			f[i] = min(f[i-coins[j]]+1, f[i])
		}
	}
	if f[amount] == VALUE {
		return -1
	}
	return f[amount]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (ans []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		ans = append(ans, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ld, rd := maxDepth(root.Left), maxDepth(root.Right)
	return max(ld, rd) + 1
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	return root
}

func lengthOfLIS(nums []int) int {
	f := make([]int, len(nums))
	for i := range f {
		f[i] = 1
	}
	for i := range nums {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				f[i] = max(f[j]+1, f[i])
			}
		}
	}
	return slices.Max(f)
}

func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for i := range nums {
		for j := target; j > 0; j-- {
			if j >= nums[i] {
				dp[j] = dp[j] || dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= nums[len(nums)-1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if target <= nums[len(nums)-1] {
		right = len(nums) - 1
	} else {
		left = 0
	}
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func maxProfit(prices []int) (ans int) {
	curMinPrice := prices[0]
	for i := range prices {
		curMinPrice = min(prices[i], curMinPrice)
		ans = max(ans, prices[i]-curMinPrice)
	}
	return
}

func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var getHeight func(node *TreeNode) int
	getHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lh, rh := getHeight(node.Left), getHeight(node.Right)
		ans = max(ans, lh+rh-2)
		return max(lh, rh) + 1
	}
	getHeight(root)
	return
}

func pathSum(root *TreeNode, targetSum int) (ans int) {
	countMap := map[int]int{}
	countMap[0] = 1
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum += node.Val
		ans += countMap[sum-targetSum]
		countMap[sum]++
		dfs(node.Left, sum)
		dfs(node.Right, sum)
		countMap[sum]--
	}
	dfs(root, 0)
	return
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	ansl, ansr := 0, 0
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] != s[j] {
				continue
			}
			if j-i == 1 {
				if j-i > ansr-ansl {
					ansr, ansl = j, i
				}
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i+1][j-1]
				if dp[i][j] {
					if ansr-ansl < j-i {
						ansr, ansl = j, i
					}
				}
			}
		}
	}
	return s[ansl : ansr+1]
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = 1 + dp[i][j]
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[m][n]
}

func singleNumber(nums []int) (ans int) {
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return
}

func dailyTemperatures(temperatures []int) []int {
	stk := []int{}
	ans := make([]int, len(temperatures))
	for i := range temperatures {
		for len(stk) > 0 && temperatures[i] > temperatures[stk[len(stk)-1]] {
			pt := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			ans[pt] = i - pt
		}
		stk = append(stk, i)
	}
	return ans
}

func maxProduct(nums []int) int {
	mx, mi := make([]int, len(nums)), make([]int, len(nums))
	mx[0], mi[0] = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx[i] = max(nums[i], nums[i]*mx[i-1], nums[i]*mi[i-1])
		mi[i] = min(nums[i], nums[i]*mx[i-1], nums[i]*mi[i-1])
	}
	return slices.Max(mx)
}

func solveNQueens(n int) (ans [][]string) {
	queenPlace := [][2]int{}
	var dfs func([][2]int, int)
	dfs = func(queenPlace [][2]int, i int) {
		if i == n {
			ans = append(ans, Plot(queenPlace, n))
			return
		}
		for col := 0; col < n; col++ {
			var flag bool
			for row := range queenPlace {
				if math.Abs(float64(i-queenPlace[row][0])) == math.Abs(float64(col-queenPlace[row][1])) || col == queenPlace[row][1] {
					flag = true
					break
				}
			}
			if !flag {
				queenPlace = append(queenPlace, [2]int{i, col})
				dfs(queenPlace, i+1)
				queenPlace = queenPlace[:len(queenPlace)-1]
			}
		}
	}
	dfs(queenPlace, 0)
	return ans
}

func Plot(queenPlace [][2]int, n int) []string {
	place := []string{}
	mrx := [][]byte{}
	for i := 0; i < n; i++ {
		t := []byte{}
		for i := 0; i < n; i++ {
			t = append(t, '.')
		}
		mrx = append(mrx, t)
	}
	for i := range queenPlace {
		mrx[queenPlace[i][0]][queenPlace[i][1]] = 'Q'
	}
	for i := range mrx {
		place = append(place, string(mrx[i]))
	}
	return place
}

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= nums[len(nums)-1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

func largestRectangleArea(heights []int) (ans int) {
	stk := []int{}
	heights = append(heights, -1)
	for i, h := range heights {
		for len(stk) > 0 && heights[stk[len(stk)-1]] > h {
			right := i
			height := heights[stk[len(stk)-1]]
			stk = stk[:len(stk)-1]
			left := -1
			if len(stk) > 0 {
				left = stk[len(stk)-1]
			}
			ans = max(ans, (right-left-1)*height)
		}
		stk = append(stk, i)
	}
	return
}

func majorityElement(nums []int) (ans int) {
	count := 0
	ans = nums[0]
	for i := range nums {
		if nums[i] == ans {
			count++
		} else {
			if count == 0 {
				ans = nums[i]
				count++
			} else {
				count--
			}
		}
	}
	return
}

func levelOrder(root *TreeNode) (ans [][]int) {
	queue := []*TreeNode{}
	layer := []int{}
	if root == nil {
		return
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		c := len(queue)
		for i := 0; i < c; i++ {
			node := queue[0]
			layer = append(layer, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, layer)
		layer = []int{}
	}
	return
}

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, lb, rb int) bool
	dfs = func(node *TreeNode, lb, rb int) bool {
		if node == nil {
			return true
		}
		val := node.Val
		return val > lb && val < rb && dfs(node.Left, lb, val) && dfs(node.Right, val, rb)
	}
	return dfs(root, math.MinInt, math.MaxInt)
}

func kthSmallest(root *TreeNode, k int) int {
	count := k
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		res := dfs(node.Left)
		if res != -1 {
			return res
		}
		count--
		if count == 0 {
			return node.Val
		}
		return dfs(node.Right)
	}
	return dfs(root)
}

func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func partition(s string) (ans [][]string) {
	path := []string{}
	var dfs func(n int)
	dfs = func(n int) {
		if n == len(s) {
			ans = append(ans, append([]string{}, path...))
			return
		}
		for i := n; i < len(s); i++ {
			if !isPalindrome(s, n, i) {
				continue
			}
			path = append(path, s[n:i+1])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}

func subsets(nums []int) (ans [][]int) {
	path := []int{}
	mp := map[int]bool{}
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}
		dfs(idx + 1)
		mp[nums[idx]] = true
		path = append(path, nums[idx])
		dfs(idx + 1)
		path = path[:len(path)-1]
		mp[nums[idx]] = false
	}
	dfs(0)
	return
}

func wordBreak(s string, wordDict []string) bool {
	mp := map[string]bool{}
	ml := 0
	for _, word := range wordDict {
		mp[word] = true
		ml = max(ml, len(word))
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(s); i++ {
		for j := i; j >= max(0, i-ml+1); j-- {
			if mp[s[j:i+1]] {
				dp[i+1] = dp[j] || dp[i+1]
				if dp[i+1] {
					break // 找到一个匹配即可
				}
			}
		}
	}
	return dp[len(s)]
}

func rightSideView(root *TreeNode) (ans []int) {
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if len(ans) == depth {
			ans = append(ans, node.Val)
		}
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}
	dfs(root, 0)
	return ans
}

func flatten(root *TreeNode) {
	var pre *TreeNode
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Right)
		postorder(node.Left)
		node.Right = pre
		node.Left = nil
	}
	postorder(root)
	return
}

func partitionLabels(s string) []int {
	lastPos := map[byte]int{}
	for i, _ := range s {
		lastPos[s[i]] = i
	}
	res := []int{}
	start, end := 0, 0
	for i := range s {
		if lastPos[s[i]] > end {
			end = lastPos[s[i]]
		}
		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res
}

func searchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := (left + right) / 2
		row, col := mid/n, mid%n
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] {
			left++
		} else {
			right--
		}
	}
	return false
}

func longestValidParentheses(s string) (ans int) {
	stk := []int{}
	for i := range s {
		if s[i] == ')' && len(stk) > 0 && s[stk[len(stk)-1]] == '(' {
			stk = stk[:len(stk)-1]
			if len(stk) == 0 {
				ans = max(ans, i+1)
			} else {
				ans = max(ans, i-stk[len(stk)-1])
			}
		} else {
			stk = append(stk, i)
		}
	}
	return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast := dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	slow := dummy
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l, r := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	return root
}

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lsum := dfs(node.Left)
		rsum := dfs(node.Right)
		ans = max(ans, lsum+rsum+node.Val)
		return max(0, max(lsum, rsum)+node.Val)
	}
	dfs(root)
	return ans
}

func decodeString(s string) string {
	stk := []string{}
	nstk := []int{}
	multi, res := 0, ""
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			multi = multi*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			stk = append(stk, res)
			nstk = append(nstk, multi)
			multi, res = 0, ""
		} else if s[i] == ']' {
			lres, n := stk[len(stk)-1], nstk[len(nstk)-1]
			stk, nstk = stk[:len(stk)-1], nstk[:len(nstk)-1]
			for i := 0; i < n; i++ {
				lres += res
			}
			res = lres
		} else {
			res += string(s[i])
		}
	}
	return res
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	slow, fast := dummy.Next, dummy.Next.Next
	for {
		slow.Val, fast.Val = fast.Val, slow.Val
		if fast.Next != nil && fast.Next.Next != nil {
			slow = fast.Next
			fast = fast.Next.Next
		} else {
			break
		}
	}
	return dummy.Next
}

func canJump(nums []int) bool {
	maxR := nums[0]
	for i := range nums {
		if i > maxR {
			return false
		}
		maxR = max(maxR, i+nums[i])
	}
	return true
}

func jump(nums []int) int {
	mxb, rb, count := 0, 0, 0
	for i := range nums[:len(nums)-1] {
		mxb = max(mxb, nums[i]+i)
		if i == rb {
			count++
			rb = mxb
		}
	}
	return count
}

func threeSum(nums []int) (ans [][]int) {
	slices.Sort(nums)
	n := len(nums)
	for k := 0; k < n-2; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		left, right := k+1, n-1
		for left < right {
			s := nums[k] + nums[left] + nums[right]
			if s == 0 {
				ans = append(ans, []int{nums[k], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if s > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return
}

func exist(board [][]byte, word string) (ans bool) {
	m, n := len(board), len(board[0])
	var dfs func(x, y, idx int)
	dir := [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, 1}, []int{0, -1}}
	dfs = func(x, y, idx int) {
		if idx == len(word) {
			ans = true
			return
		}
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if word[idx] != board[x][y] || board[x][y] == '0' {
			return
		}
		for i := range dir {
			nx, ny := x+dir[i][0], y+dir[i][1]
			board[x][y] = byte('0')
			dfs(nx, ny, idx+1)
			board[x][y] = word[idx]
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, 0)
			if ans {
				return
			}
		}
	}
	return false
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy, n := &ListNode{Next: head}, 0
	p, p0 := dummy, dummy
	for p.Next != nil {
		n++
		p = p.Next
	}
	cur := p0.Next
	var pre, nxt *ListNode
	for ; n >= k; n -= k {
		for i := 0; i < k; i++ {
			nxt = cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}

		temp := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = temp
	}
	return dummy.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	for list1 != nil {
		cur.Next = list1
		list1 = list1.Next
		cur = cur.Next
	}
	for list2 != nil {
		cur.Next = list2
		list2 = list2.Next
		cur = cur.Next
	}
	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var partition func(lists []*ListNode, left, right int) *ListNode
	partition = func(lists []*ListNode, left, right int) *ListNode {
		if right == left {
			return lists[left]
		}
		if right == left+1 {
			return mergeTwoLists(lists[left], lists[right])
		}
		mid := left + (right-left)/2
		ll := partition(lists, left, mid)
		rl := partition(lists, mid+1, right)
		return mergeTwoLists(ll, rl)
	}
	return partition(lists, 0, len(lists)-1)
}

func maxProfit3(prices []int) int {
	n := len(prices)
	dp := make([][][2]int, n+1)
	for i := range dp {
		dp[i] = make([][2]int, 4)
		for j := range dp[i] {
			dp[i][j] = [2]int{math.MinInt / 2, math.MinInt / 2}
		}
	}
	for i := 1; i <= 3; i++ {
		dp[0][i][0] = 0
	}
	for i := range prices {
		for k := 1; k <= 3; k++ {
			dp[i+1][k][0] = max(dp[i][k][1]+prices[i], dp[i][k][0])
			dp[i+1][k][1] = max(dp[i][k-1][0]-prices[i], dp[i][k][1])
		}
	}
	return dp[n][3][0]
}

func canPartitionKSubsets(nums []int, k int) (ans bool) {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	sort.Ints(nums)
	used := make([]bool, len(nums))
	target := sum / k
	n := len(nums)
	if nums[0] > target || nums[n-1] > target {
		return false
	}
	var dfs func(idx int, count int, curSum int) bool
	dfs = func(idx int, count int, curSum int) bool {
		if count == k-1 {
			return true
		}
		if curSum == target {
			return dfs(0, count+1, 0)
		}
		for i := idx; i < n; i++ {
			if used[i] || curSum+nums[i] > target {
				continue
			}
			used[i] = true
			res := dfs(idx+1, count, curSum+nums[i])
			if res {
				return true
			}
			used[i] = false
			if curSum == 0 {
				return false
			}

		}
		return false
	}

	return dfs(0, 0, 0)
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	nodeMap := map[*Node]*Node{}
	cur := head
	for cur != nil {
		newNode := &Node{Val: cur.Val}
		nodeMap[cur] = newNode
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		newNode := nodeMap[cur]
		newNode.Next = nodeMap[cur.Next]
		newNode.Random = nodeMap[cur.Random]
		cur = cur.Next
	}
	return nodeMap[head]
}

func combinationSum(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	if candidates[0] > target {
		return
	}
	path := []int{}
	var dfs func(idx, sum int)
	dfs = func(idx, sum int) {
		if sum == target {
			ans = append(ans, append([]int{}, path...))
			return
		}
		if candidates[idx]+sum > target {
			return
		}
		path = append(path, candidates[idx])
		dfs(idx, sum+candidates[idx])
		path = path[:len(path)-1]
		if idx+1 < len(candidates) {
			dfs(idx+1, sum)
		}
	}
	dfs(0, 0)
	return
}

type LRUCache struct {
	kvMap map[int]*dNode
	cap   int
	len   int
	dummy *dNode
}

type dNode struct {
	Key  int
	Val  int
	pre  *dNode
	next *dNode
}

func Constructor(capacity int) LRUCache {
	dummy := &dNode{} // 先创建dummy节点
	dummy.pre, dummy.next = dummy, dummy
	lru := LRUCache{cap: capacity, kvMap: map[int]*dNode{}, dummy: dummy}
	return lru
}

func (cache *LRUCache) Get(key int) int {
	node, _ := cache.kvMap[key]
	if node == nil {
		return -1
	}

	if node.pre == cache.dummy {
		return node.Val
	}

	node.pre.next = node.next
	node.next.pre = node.pre

	node.next = cache.dummy.next
	cache.dummy.next.pre = node
	cache.dummy.next = node
	node.pre = cache.dummy

	return node.Val
}

func (cache *LRUCache) Put(key int, value int) {
	if cache.kvMap[key] != nil {
		node := cache.kvMap[key]
		node.Val = value
		node.pre.next = node.next
		node.next.pre = node.pre

		node.next = cache.dummy.next
		cache.dummy.next.pre = node
		cache.dummy.next = node
		node.pre = cache.dummy
		return
	}

	newNode := &dNode{Val: value, Key: key}
	cache.kvMap[key] = newNode
	if cache.len < cache.cap {
		cache.len++
	} else {
		lastNode := cache.dummy.pre
		cache.dummy.pre = lastNode.pre
		lastNode.pre.next = lastNode.next
		delete(cache.kvMap, lastNode.Key)
	}
	newNode.next = cache.dummy.next
	cache.dummy.next.pre = newNode
	cache.dummy.next = newNode
	newNode.pre = cache.dummy
}

func isValidBST2(root *TreeNode) bool {
	pre := math.MinInt
	var inorder func(node *TreeNode) bool
	inorder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		lres := inorder(node.Left)
		if !lres || node.Val <= pre {
			return false
		}
		pre = node.Val
		return inorder(node.Right)
	}
	return inorder(root)
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}
	for i := range word1 {
		for j := range word2 {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1]) + 1
			}
		}
	}
	return dp[m][n]
}

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (hp *Heap) Pop() any {
	h := *hp
	x := h[len(h)-1]
	h = h[:len(h)-1]
	return x
}

func (hp *Heap) Push(x any) {
	h := *hp
	h = append(h, x.(int))
	*hp = h
}

func findKthLargest(nums []int, k int) int {
	hp := &Heap{}
	for i := range nums {
		heap.Push(hp, nums[i])
		if hp.Len() > k {
			heap.Pop(hp)
		}
	}

	x := heap.Pop(hp).(int)
	return x
}

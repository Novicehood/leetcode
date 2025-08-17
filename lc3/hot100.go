package main

import (
	"fmt"
	"maps"
	"math"
	"slices"
)

func moveZeroes(nums []int) {
	i0 := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[i0], nums[i] = nums[i], nums[i0]
			i0++
		}
	}
	return
}

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for i := range strs {
		bytes := []byte(strs[i])
		slices.Sort(bytes)
		s := string(bytes)
		if mp[s] == nil {
			mp[s] = []string{}
		}
		mp[s] = append(mp[s], strs[i])
	}
	return slices.Collect(maps.Values(mp))
}

func threeSum(nums []int) (ans [][]int) {
	slices.Sort(nums)
	n := len(nums)
	for k := 0; k < n-2; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		l, r := k+1, n-1
		for l < r {
			sum := nums[k] + nums[l] + nums[r]
			if sum == 0 {
				ans = append(ans, []int{nums[k], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return
}

func lengthOfLongestSubstring(s string) (ans int) {
	mp := map[byte]int{}
	l := 0
	for r := range s {
		mp[s[r]]++
		for mp[s[r]] > 1 {
			mp[s[l]]--
			l++
		}
		ans = max(ans, r-l+1)
	}
	return
}

func findAnagrams(s string, p string) (res []int) {
	pmp, smp := map[byte]int{}, map[byte]int{}
	for i := range p {
		pmp[p[i]]++
	}
	l := 0
	for i := range s {
		smp[s[i]]++
		for smp[s[i]] > pmp[s[i]] {
			smp[s[l]]--
			l++
		}
		if i-l+1 == len(p) {
			res = append(res, l)
		}
	}
	return
}

func trap(height []int) (ans int) {
	n := len(height)
	lmx, rmx, l, r := make([]int, n), make([]int, n), 0, 0
	for i := 0; i < n; i++ {
		if i == 0 {
			lmx[i], rmx[n-1] = height[i], height[n-1]
			l, r = height[0], height[n-1]
		} else {
			l, r = max(l, height[i]), max(r, height[n-i-1])
			lmx[i], rmx[n-1-i] = l, r
		}
	}
	for i := range height {
		ans += min(lmx[i], rmx[i]) - height[i]
	}
	return
}

func subarraySum(nums []int, k int) (ans int) {
	mp := map[int]int{}
	mp[0] = 1
	sum := 0
	for i := range nums {
		sum += nums[i]
		ans += mp[sum-k]
		mp[sum]++
	}
	return
}

func maxSlidingWindow(nums []int, k int) (ans []int) {
	q := []int{}
	for i := range nums {
		for len(q) > 0 && nums[q[len(q)-1]] <= nums[i] {
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

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		for nums[i] >= 1 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			dp[0] = nums[0]
		} else {
			dp[i] = max(dp[i-1]+nums[i], nums[i])
		}
	}
	return slices.Max(dp)
}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	row, col := 0, n-1
	for row < m && col >= 0 {
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var pre *TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		dfs(node.Left)
		node.Right = pre
		node.Left = nil
		pre = node
	}
	dfs(root)
}

func pathSum(root *TreeNode, targetSum int) (ans int) {
	mp := map[int]int{}
	mp[0] = 1
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum += node.Val
		ans += mp[sum-targetSum]
		mp[sum]++
		dfs(node.Left, sum)
		dfs(node.Right, sum)
		mp[sum]--
	}
	dfs(root, 0)
	return
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	node0, node1 := dummy, dummy.Next
	for node1 != nil && node1.Next != nil {
		node2 := node1.Next
		node3 := node2.Next

		node0.Next = node2
		node2.Next = node1
		node1.Next = node3

		node0 = node1
		node1 = node3
	}
	return dummy.Next
}

func twoSum(nums []int, target int) []int {
	mp := map[int]int{}
	for i := range nums {
		if v, ok := mp[target-nums[i]]; ok {
			return []int{i, v}
		}
		mp[nums[i]] = i
	}
	return []int{-1, -1}
}

func maxArea(height []int) (ans int) {
	l, r := 0, len(height)-1
	for l < r {
		ans = max(ans, (r-l)*min(height[l], height[r]))
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return
}

func minWindow(s string, t string) string {
	mp := map[byte]int{}
	remain := 0
	for i := range t {
		mp[t[i]]++
		if mp[t[i]] == 1 {
			remain++
		}
	}
	l := 0
	ansl, ansr := 0, len(s)
	for r := range s {
		mp[s[r]]--
		if mp[s[r]] == 0 {
			remain--
		}
		for remain == 0 {
			if r-l < ansr-ansl {
				ansr, ansl = r, l
			}
			mp[s[l]]++
			if mp[s[l]] == 1 {
				remain++
			}
			l++
		}
	}
	if ansr-ansl == len(s) {
		return ""
	}
	return s[ansl : ansr+1]
}

func detectCycle(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil || slow == nil {
		return nil
	}
	fast = dummy
	for fast != slow {
		fast, slow = fast.Next, slow.Next
	}
	return fast
}

func generateParenthesis(n int) (ans []string) {
	var dfs func(idx, countl int)
	path := []byte{}
	dfs = func(idx, countl int) {
		if idx == 2*n {
			ans = append(ans, string(path))
			return
		}
		if countl < n {
			path = append(path, '(')
			dfs(idx+1, countl+1)
			path = path[:len(path)-1]
		}
		if idx-countl < countl {
			path = append(path, ')')
			dfs(idx+1, countl)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return
}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

type DNode struct {
	pre   *DNode
	next  *DNode
	key   int
	value int
}

type LRUCache struct {
	sentinel *DNode
	hm       map[int]*DNode
	len      int
	cap      int
}

func Constructor(capacity int) LRUCache {
	lrc := LRUCache{
		sentinel: &DNode{},
		hm:       map[int]*DNode{},
		len:      0,
		cap:      capacity,
	}
	lrc.sentinel.pre = lrc.sentinel
	lrc.sentinel.next = lrc.sentinel
	return lrc
}

func (this *LRUCache) Get(key int) int {
	vd, ok := this.hm[key]
	if !ok {
		return -1
	}
	//	更新
	if this.len > 1 {
		if vd.pre != this.sentinel {
			vd.pre.next = vd.next
			vd.next.pre = vd.pre

			vd.next = this.sentinel.next
			this.sentinel.next.pre = vd
			this.sentinel.next = vd
			vd.pre = this.sentinel
		}
	}
	return vd.value
}

func (this *LRUCache) Put(key int, value int) {
	vd, ok := this.hm[key]
	if ok {
		//	存在
		vd.value = value
		//	移动
		if vd.pre != this.sentinel {
			vd.pre.next = vd.next
			vd.next.pre = vd.pre

			vd.next = this.sentinel.next
			this.sentinel.next.pre = vd
			this.sentinel.next = vd
			vd.pre = this.sentinel
		}
		return
	}
	//	需要插入，先插入
	newnode := &DNode{
		key:   key,
		value: value,
	}
	this.hm[key] = newnode
	newnode.next = this.sentinel.next
	this.sentinel.next.pre = newnode
	newnode.pre = this.sentinel
	this.sentinel.next = newnode

	if this.len == this.cap {
		//	需要删除
		k := this.sentinel.pre.key
		delete(this.hm, k)
		this.sentinel.pre.pre.next = this.sentinel
		this.sentinel.pre = this.sentinel.pre.pre
		return
	}
	this.len++
	return
}

func maxPathSum(root *TreeNode) (ans int) {
	ans = math.MinInt / 2
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lsum := dfs(node.Left)
		rsum := dfs(node.Right)
		ans = max(ans, node.Val+lsum+rsum)
		return max(0, max(lsum, rsum)+node.Val)
	}
	dfs(root)
	return
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

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	n := len(nums)
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= nums[n-1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if target <= nums[n-1] {
		right = n - 1
	} else {
		left = 0
	}
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	var lowerBound func(target int) int
	lowerBound = func(target int) int {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := (left + right) / 2
			if nums[mid] >= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return left
	}
	res1, res2 := lowerBound(target), lowerBound(target+1)
	if res1 >= len(nums) || nums[res1] != target {
		return []int{-1, -1}
	}
	return []int{res1, res2}
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
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[target]
}

func kthSmallest(root *TreeNode, k int) int {
	var inorder func(node *TreeNode) int
	inorder = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		lres := inorder(node.Left)
		if lres != -1 {
			return lres
		}
		k--
		if k == 0 {
			return node.Val
		}
		return inorder(node.Right)
	}
	return inorder(root)
}

func swapPairs2(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := head
	pre := dummy
	for cur != nil && cur.Next != nil {
		nxt := cur.Next.Next
		n0 := cur
		n1 := cur.Next

		n1.Next = n0
		n0.Next = nxt
		pre.Next = n1
		pre = n0

		cur = nxt
	}
	return dummy.Next
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var getHeight func(node *TreeNode) int
	getHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lh, rh := getHeight(node.Left), getHeight(node.Right)
		ans = max(ans, lh+rh)
		return max(lh, rh) + 1
	}
	getHeight(root)
	return
}

func maxPathSum2(root *TreeNode) (ans int) {
	var getSubSum func(node *TreeNode) int
	getSubSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lsum, rsum := getSubSum(node.Left), getSubSum(node.Right)
		ans = max(ans, node.Val+lsum+rsum)
		return max(0, max(lsum, rsum)+node.Val)
	}
	getSubSum(root)
	return
}

func numIslands(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x > m || y < 0 || y > n || grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return
}

func MaxValueLessN(n int) {
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	x := n
	ns := []int{}
	for x > 0 {
		ns = append(ns, x%10)
		x = x / 10
	}
	ln := len(ns)
	slices.Reverse(ns)
	mx := 0
	var dfs func(idx, sum int, match bool)
	dfs = func(idx, sum int, match bool) {
		if idx == ln {
			if sum < n {
				mx = max(mx, sum)
			}
			return
		}
		for i := range nums {
			if nums[i] == ns[idx] && match {
				dfs(idx+1, sum*10+nums[i], match)
			}
			if !match && sum*10+nums[i] < n {
				dfs(idx+1, sum*10+nums[i], false)
			}
		}
	}
	dfs(0, 0, false)
	fmt.Println(mx)
}

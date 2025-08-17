package lc4

import (
	"maps"
	"slices"
	"sort"
)

func longestConsecutive(nums []int) (ans int) {
	mp := map[int]bool{}
	for i := range nums {
		mp[nums[i]] = true
	}
	for k, _ := range mp {
		if mp[k-1] {
			continue
		}
		y := k
		for mp[y] {
			y++
		}
		ans = max(ans, k-y)
	}
	return
}

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for i := range strs {
		strb := []byte(strs[i])
		slices.Sort(strb)
		sortStr := string(strb)
		if mp[sortStr] == nil {
			mp[sortStr] = []string{}
		}
		mp[sortStr] = append(mp[sortStr], strs[i])
	}
	res := slices.Collect(maps.Values(mp))
	return res
}

func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1
		if nums[i] > 0 || i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{nums[left], nums[right], nums[i]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return
}

func moveZeroes(nums []int) {
	var i0 int
	for i := range nums {
		if nums[i] != 0 {
			nums[i0], nums[i] = nums[i], nums[i0]
			i0++
		}
	}
}

func sortColors(nums []int) {
	var ired int
	for i := range nums {
		if nums[i] == 0 {
			nums[i], nums[ired] = nums[ired], nums[i]
			ired++
		}
	}
	iblue := len(nums) - 1
	for i := len(nums) - 1; i >= ired; i-- {
		if nums[i] == 2 {
			nums[iblue], nums[i] = nums[i], nums[iblue]
			iblue--
		}
	}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	indegree := make([]int, numCourses)
	depnd := map[int][]int{}
	for i := range prerequisites {
		c1, c2 := prerequisites[i][0], prerequisites[i][1]
		indegree[c1]++
		if _, ok := depnd[c2]; !ok {
			depnd[c2] = []int{}
		}
		depnd[c2] = append(depnd[c2], c1)
	}
	q := []int{}
	count := 0
	for i := range indegree {
		if indegree[i] == 0 {
			q = append(q, i)
			count++
		}
	}
	for len(q) > 0 {
		c := q[0]
		q = q[1:len(q)]
		clist := depnd[c]
		for i := range clist {
			indegree[clist[i]]--
			if indegree[clist[i]] == 0 {
				q = append(q, clist[i])
				count++
			}
		}
	}
	if count == numCourses {
		return true
	}
	return false
}

func lengthOfLongestSubstring(s string) (ans int) {
	mp, l := map[byte]int{}, 0
	for i := range s {
		mp[s[i]]++
		for mp[s[i]] > 1 {
			mp[s[l]]--
			l++
		}
		ans = max(ans, i-l+1)
	}
	return
}

func findAnagrams(s string, p string) (ans []int) {
	mp := map[byte]int{}
	for i := range p {
		mp[p[i]]++
	}
	var l int
	for i := range s {
		mp[s[i]]--
		for mp[s[i]] < 0 {
			mp[s[l]]++
			l++
		}
		if i-l+1 == len(p) {
			ans = append(ans, l)
		}
	}
	return
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := range text1 {
		for j := range text2 {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
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

func minWindow(s string, t string) string {
	mp := map[byte]int{}
	kind := 0
	for i := range t {
		mp[t[i]]++
		if mp[t[i]] == 1 {
			kind++
		}
	}
	ansl, ansr, l := 0, len(s), 0
	for r := range s {
		mp[s[r]]--
		if mp[s[r]] == 0 {
			kind--
		}
		for kind == 0 {
			if r-l < ansr-ansl {
				ansr, ansl = r, l
			}
			mp[s[l]]++
			if mp[s[l]] == 1 {
				kind++
			}
			l++
		}
	}
	if ansr-ansl == len(s) {
		return ""
	}
	return s[ansl : ansr+1]
}

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	for i := range nums {
		dp[i+1] = max(dp[i]+nums[i], nums[i])
	}
	return slices.Max(dp[1:])
}

func merge(intervals [][]int) (ans [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	start, end := intervals[0][0], intervals[0][1]
	for i := range intervals {
		if intervals[i][0] <= end {
			end = max(end, intervals[i][1])
		} else {
			ans = append(ans, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		}
	}
	ans = append(ans, []int{start, end})
	return
}

func firstMissingPositive(nums []int) int {
	for i := range nums {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func maxSlidingWindow(nums []int, k int) (ans []int) {
	q := []int{}
	for i := range nums {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i-q[0] >= k {
			q = q[1:]
		}
		if i >= k-1 {
			ans = append(ans, nums[q[0]])
		}
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	cur, n := head, 0
	for cur != nil {
		n++
		cur = cur.Next
	}
	p0 := dummy
	cur = head
	var pre *ListNode
	for ; n >= k; n -= k {
		for i := 0; i < k; i++ {
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}
		temp := p0.Next
		p0.Next = pre
		temp.Next = cur
		p0 = temp
	}
	return dummy.Next
}

func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= nums[n-1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

func isValid(s string) bool {
	stk := []byte{}
	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stk = append(stk, s[i])
		} else {
			if len(stk) == 0 {
				return false
			}
			if s[i] == ']' && stk[len(stk)-1] != '[' || s[i] == ')' && stk[len(stk)-1] != '(' || s[i] == '}' && stk[len(stk)-1] != '{' {
				return false
			} else {
				stk = stk[:len(stk)-1]
			}
		}
	}
	if len(stk) == 0 {
		return true
	}
	return false
}

type LRUCache struct {
	kvMap map[int]*DNode
	dummy *DNode
	cap   int
	len   int
}

type DNode struct {
	Key  int
	Val  int
	Pre  *DNode
	Next *DNode
}

func Constructor(capacity int) LRUCache {
	dummy := &DNode{}
	dummy.Pre, dummy.Next = dummy, dummy
	cache := LRUCache{kvMap: map[int]*DNode{}, cap: capacity, dummy: dummy}
	return cache
}

func (this *LRUCache) Get(key int) (ans int) {
	node, ok := this.kvMap[key]
	if !ok {
		return -1
	}
	ans = node.Val
	if node.Pre == this.dummy {
		return
	}
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	node.Next = this.dummy.Next
	this.dummy.Next.Pre = node
	this.dummy.Next = node
	node.Pre = this.dummy
	return
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.kvMap[key]
	if ok {
		node.Val = value
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
		node.Next = this.dummy.Next
		this.dummy.Next.Pre = node
		this.dummy.Next = node
		node.Pre = this.dummy
		return
	}

	newNode := &DNode{Key: key, Val: value}
	this.kvMap[key] = newNode
	if this.len < this.cap {
		this.len++
	} else {
		lastNode := this.dummy.Pre
		delete(this.kvMap, lastNode.Key)
		lastNode.Pre.Next = lastNode.Next
		lastNode.Next.Pre = lastNode.Pre
	}
	newNode.Next = this.dummy.Next
	this.dummy.Next.Pre = newNode
	this.dummy.Next = newNode
	newNode.Pre = this.dummy
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast, slow = fast.Next, slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

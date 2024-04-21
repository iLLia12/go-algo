package main

import (
	"fmt"
	"math"
	"strconv"
)

type NodeList struct {
	key int
	value int
	next *NodeList
	prev *NodeList
}

type LRUCache struct {
	capacity int
	cache map[int]*NodeList
	head *NodeList
	tail *NodeList
}

func Constructor(capacity int) LRUCache {
	head := &NodeList{0,0, nil, nil}
	tail := &NodeList{0,0, nil, nil}
	head.next = tail
	tail.prev = head
	return LRUCache{
		cache: make(map[int]*NodeList),
		capacity: capacity,
		head: head,
		tail: tail,
	}
}

func (this *LRUCache) Insert(node *NodeList) {
	prev := this.tail.prev
	prev.next = node
	node.prev = prev
	node.next = this.tail
	this.tail.prev = node
}

func (this *LRUCache) Delete(node *NodeList) {
	prev, next := node.prev, node.next
	prev.next, next.prev = next, prev
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if ok {
		this.Delete(node)
		this.Insert(node)
		return this.cache[key].value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
			node.value = value
			this.Delete(node)
			this.Insert(node)
			return
	}

	newNode := &NodeList{key: key, value: value}
	this.cache[key] = newNode
	this.Insert(newNode)

	if len(this.cache) > this.capacity {
			lru := this.head.next
			this.Delete(lru)
			delete(this.cache, lru.key)
	}
}

func exist(board [][]byte, word string) bool {
	path := make(map[string]bool)
	ROWS := len(board)
	COLS := len(board[0])
	var dfs func(c,r,i int) bool
	dfs = func(c, r, i int) bool {
		if i == len(word) { return true }
		key := fmt.Sprintf("%d-%d", r, c)
		_, ok := path[key]; 
		if c < 0 || r < 0 || r >= ROWS || c >= COLS || word[i] != board[r][c] || ok {
				return false
			}
		path[key] = true
		res := dfs(r + 1, c, i + 1) ||
				dfs(r - 1, c, i + 1) ||
				dfs(r, c + 1, i + 1) ||
				dfs(r, c - 1, i + 1)
		delete(path, key)
		return res
	}
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if dfs(r, c, 0) { return true }
		}
	}
	return false
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" && num2 == "0" {
		return "0"
	}
	n1, n2 := len(num1), len(num2)
	result := make([]int, n1 + n2)

	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			mul := (num1[i] - '0') * (num2[j] - '0')
			p1, p2 := i + j, i + j + 1
			sum := int(mul) + result[p2]

			result[p1] += sum / 10
			result[p2] = sum % 10 
		}
	}

	ans := ""
	for _, val := range result {
		if !(len(ans) == 0 && val == 0) {
			ans += strconv.Itoa(val)
		}
	}
  return ans
}

func containsNearbyDuplicate(nums []int, k int) bool {
	window := map[int]struct{}{}
	left := 0
	for right, value := range nums {
		if right - left > k {
			delete(window, value)
			left++
		}
		if _, ok := window[value]; ok {
			return true
		} else {
			window[value] = struct{}{}
		}
	}
	return false
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	if n == 0 { return }
	i := m + n - 1
	for n != 0 {
		if m > 0 && nums1[m - 1] > nums2[n - 1] {
			nums1[i] = nums1[m - 1]
			m--
		} else {
			nums1[i] = nums2[n - 1]
			n-- 
		}
		i--
	}
	fmt.Println(nums1)
}

func maxProfit(prices []int) int {
	if len(prices) < 2 { return 0 }
	left := prices[0]
	profit := 0.0
	for _, right := range prices {
		profit = math.Max(profit, float64(right - left))
		if left > right {
			left = right
		}
	}
	return int(profit)
}

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	res := ""
	carry := 0
	for i >= 0 || j >= 0 || carry != 0 {
		total := carry
		if i >= 0 {
			total += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			total += int(b[j] - '0')
			j--
		}
		res = fmt.Sprintf("%d%s", total%2, res)
		carry = total / 2
	}
	return res
}

func main() {
	
}
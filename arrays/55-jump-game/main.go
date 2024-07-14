package jumpgame

import "fmt"

func max(a int, b int) int {
	if a > b { return a }
	return b
}

func canJump(nums []int) bool {
	m := 0
l := len(nums)
for i := range len(nums) {
	m = max(m, i + nums[i])
	if m >= l - 1 { return true }
	if 0 <= i - m { return false }
}
return true
}

func main() {
	a := []int{3,2,1,0,4}
	b := []int{2,3,1,1,4}
	c := []int{0}
  fmt.Println(canJump(a))
	fmt.Println(canJump(b))
	fmt.Println(canJump(c))
}
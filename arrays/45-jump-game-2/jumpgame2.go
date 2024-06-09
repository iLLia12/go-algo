package jumpgame2

func jump(nums []int) int {
	var max func(a int, b int) int
	max = func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	count := 0
	farthest := 0
	currentEnd := 0
	for i := 0; i < len(nums) - 1; i++ {
		farthest = max(farthest, i + nums[i])
		if i == currentEnd {
			count++
			currentEnd = farthest
		}
		if currentEnd >= len(nums) - 1 {
			break
		}
	}
	return count
}
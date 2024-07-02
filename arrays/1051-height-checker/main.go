package heightchecker

import "sort"

func heightChecker(heights []int) int {
	res := 0
	heightsCopy := make([]int, len(heights))
	copy(heightsCopy, heights)
	sort.Ints(heightsCopy)
	for i := range heights {
		if heightsCopy[i] != heights[i] {
			res++
		}
	}
	return res
}
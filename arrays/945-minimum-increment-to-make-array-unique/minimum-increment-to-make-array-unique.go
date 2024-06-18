package minimumincrementtomakearrayunique

import "sort"

func minIncrementForUnique(nums []int) int {
	res := 0
	sort.Ints(nums)
	for i := 1; len(nums) > i; i++ {
		if nums[i - 1] == nums[i] {
			res++
			nums[i]++
		} else if nums[i - 1] > nums[i] {
			a := nums[i - 1] - nums[i] + 1
			nums[i] += a
			res += a
		}
	}
	return res
}
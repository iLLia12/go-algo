package minimumnumberofmovestoseateveryone

import "sort"

func minMovesToSeat(seats []int, students []int) int {
	res := 0
	sort.Ints(seats)
	sort.Ints(students)
	i := 0
	for len(seats) > i {
		dif := seats[i] - students[i]
		if dif < 0 { 
			res += -dif 
		} else { res += dif }
		i++
	}
	return res
}
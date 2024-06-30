package largestlocalvaluesinamatrix


func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}


func largestLocal(grid [][]int) [][]int {
	var res [][]int
	i := 1
	for i < len(grid) - 1 {
		j := 1
		var n []int
		for j < len(grid) - 1 {
			i1 := 0
			m := 0
			for i1 < 3 {
				j1 := 0
				for j1 < 3 {
					m = max(m, grid[i - 1 + i1][j - 1 + j1])
					j1++
				}
				i1++
			}
			n = append(n, m)
			j++
		}
		res = append(res, n)
		i++
	}
	return res
}
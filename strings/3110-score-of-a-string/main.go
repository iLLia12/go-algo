package scoreofastring

import "math"

func scoreOfString(s string) int {
	res := 0
	for i := 1; i < len(s); i++ {
		res += int(math.Abs(float64(int(s[i - 1]) - int(s[i]))))
	}
	return res
}
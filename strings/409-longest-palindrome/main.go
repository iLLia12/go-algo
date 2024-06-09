package main

func longestPalindrome(s string) int {
	m := make(map[byte]int)
	count := 0
	for _, char := range s {
		val, ok := m[byte(char)]
		if !ok || val == 0 {
			m[byte(char)] = 1
		} else {
			m[byte(char)] = 0
			count++
		}
	}
	res := count * 2
	if len(s) - res != 0 {
		res++
	}
	return res
}
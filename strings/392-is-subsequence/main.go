package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 { return true }
	j := -1
	for i := range t {
		if t[i] == s[j + 1] {
			j++
		}
		if j == len(s) - 1 {
			return true
		}
	}
	return false
}
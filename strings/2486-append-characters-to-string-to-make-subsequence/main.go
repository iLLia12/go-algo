package appendcharacterstostringtomakesubsequence

func appendCharacters(s string, t string) int {
	j := 0
	for i := range s {
		if(j < len(t) && t[j] == s[i]) { j++ }
	}
	return len(t) - j
}
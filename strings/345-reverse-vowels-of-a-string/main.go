package main

func inSlice(slice []byte, search byte) bool {
	for _, val := range slice {
		if val == search {
			return true
		}
	}
	return false
}

func reverseVowels(s string) string {
	vowels := []byte{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}
	left := 0
	right := len(s) - 1
	chars := []byte(s)
	for left < right {
		if inSlice(vowels, chars[left]) {
			for left < right {
				if inSlice(vowels, chars[right]) {
					chars[left], chars[right] = chars[right], chars[left]
					right--
					break
				} else { right-- }
			}
		}
		left++
	}
	return string(chars)
}
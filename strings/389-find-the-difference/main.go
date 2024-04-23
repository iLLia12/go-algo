package main

func findTheDifference(s string, t string) byte {
	m := make(map[byte]int)
	for _, char := range t {
		_, ok := m[byte(char)]
		if !ok {
			m[byte(char)] = 1
		} else {
			m[byte(char)] += 1
		}
	}
	for _, char := range s {
		_, ok := m[byte(char)]
		if ok {
			m[byte(char)] -= 1
		}
	}

	for key, value := range m {
		if value != 0 {
			return key
		}
	}

	return 0
}
package Solution

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var xor, squaresum1, squaresum2 rune
	for _, char := range s {
		xor ^= char
		squaresum1 += char * char
	}
	for _, char := range t {
		xor ^= char
		squaresum2 += char * char
	}
	return xor == 0 && squaresum1 == squaresum2
}

func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var count [26]int
	for _, c := range s {
		count[c-'a']++
	}

	for _, c := range t {
		count[c-'a']--
		if count[c-'a'] < 0 {
			return false
		}
	}
	return true
}

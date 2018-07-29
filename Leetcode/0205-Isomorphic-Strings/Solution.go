package Solution

func isIsomorphic(s string, t string) bool {
	s1 := [256]int{}
	t1 := [256]int{}
	for i := 0; i < len(s); i++ {
		sc := s[i]
		tc := t[i]
		if s1[sc] != t1[tc] {
			return false
		}
		s1[sc] = i + 1
		t1[tc] = i + 1
	}
	return true
}

func isIsomorphic1(s string, t string) bool {
	m1 := map[byte]int{}
	m2 := map[byte]int{}
	for i := 0; i < len(s); i++ {
		sKey := s[i]
		tKey := t[i]
		if m1[sKey] != m2[tKey] {
			return false
		}
		m1[sKey] = i + 1
		m2[tKey] = i + 1
	}
	return true
}

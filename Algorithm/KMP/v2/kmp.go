package v2

const (
	//MaxArraySize Need modify this value if your pattern is much learger than 100
	MaxArraySize int = 100
)

//Strchr :Returns a pointer to the last occurrence of string in the  string str
func Strchr(haystack string, needle string) int {
	retSlice := KMP(haystack, needle)
	if len(retSlice) > 0 {
		return retSlice[len(retSlice)-1]
	}

	return -1
}

//Strstr :Use kmp for strstr function
func Strstr(haystack string, needle string) int {
	retSlice := KMP(haystack, needle)
	if len(retSlice) > 0 {
		return retSlice[0]
	}

	return -1
}

//KMP Return index list if target string contain want string
func KMP(haystack string, needle string) []int {
	next := preKMP(needle)
	i := 0
	j := 0
	m := len(needle)
	n := len(haystack)

	x := []byte(needle)
	y := []byte(haystack)
	var ret []int

	//got zero target or want, just return empty result
	if m == 0 || n == 0 {
		return ret
	}

	//want string bigger than target string
	if n < m {
		return ret
	}

	for j < n {
		for i > -1 && x[i] != y[j] {
			i = next[i]
		}
		i++
		j++

		//fmt.Println(i, j)
		if i >= m {
			ret = append(ret, j-i)
			//fmt.Println("find:", j, i)
			i = next[i]
		}
	}

	return ret
}

func preMP(x string) [MaxArraySize]int {
	var i, j int
	length := len(x) - 1
	var mpNext [MaxArraySize]int
	i = 0
	j = -1
	mpNext[0] = -1

	for i < length {
		for j > -1 && x[i] != x[j] {
			j = mpNext[j]
		}
		i++
		j++
		mpNext[i] = j
	}
	return mpNext
}

func preKMP(x string) [MaxArraySize]int {
	var i, j int
	length := len(x) - 1
	var kmpNext [MaxArraySize]int
	i = 0
	j = -1
	kmpNext[0] = -1

	for i < length {
		for j > -1 && x[i] != x[j] {
			j = kmpNext[j]
		}

		i++
		j++

		if x[i] == x[j] {
			kmpNext[i] = kmpNext[j]
		} else {
			kmpNext[i] = j
		}
	}
	return kmpNext
}

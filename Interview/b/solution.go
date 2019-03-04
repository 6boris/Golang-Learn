package b

/*
	s[i-1] s[i] s[i+1]

	abb   aeb
	aaa   aea
	baa	  bea
*/
func Solution(str string) int {

	num := 0
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			num++
			i++
		}
	}
	return 1
}

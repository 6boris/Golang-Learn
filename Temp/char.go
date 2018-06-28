package main

import "fmt"

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func main() {
	str1 := "asadadsdadsd"
	str2 := "asd"

	fmt.Println(HasPrefix(str1, str2))
}

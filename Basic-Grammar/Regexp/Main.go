package main

import (
	"fmt"
	"html/template"
	"regexp"
)

func main() {
	m, err := regexp.MatchString("^[0-9]+$", "9s9")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
	fmt.Println("username:", template.HTMLEscapeString("<script>alert('you have been pwned')</script>"))
}

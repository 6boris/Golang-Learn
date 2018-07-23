package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.Contains("in failure", "a "))
	fmt.Println(strings.ContainsAny("in failure", "s g"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))

	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
	fmt.Println(strings.FieldsFunc("  foo bar  baz   \n", unicode.IsSpace))
}

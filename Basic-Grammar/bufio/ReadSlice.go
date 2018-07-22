package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.Contains("in failure", "a "))
	fmt.Println(strings.ContainsAny("in failure", "s g"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
	fmt.Println(reflect.TypeOf(1 > 2))
	fmt.Println(strings.Count("asdasda", "as"))
	fmt.Println(strings.Count("fivevev", "vev"))
}

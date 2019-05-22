package main

import (
	"fmt"
	"log"
	"plugin"
)

func main() {
	p, err := plugin.Open("str.so")
	if err != nil {
		log.Panicf("plugin.Open: %s\n", err)
	}
	f, err := p.Lookup("UpperCase")
	if err != nil {
		log.Panicf("Lookup UpperCase: %s\n", err)
	}
	UpperCase, ok := f.(func(string) string)
	if !ok {
		log.Panicf("UpperCase assertion: %s\n", err)
	}
	s := UpperCase("hello")
	fmt.Println(s)
}

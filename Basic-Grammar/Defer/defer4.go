package main

import (
	"errors"
	"fmt"
)

func e1() {
	var err error
	defer fmt.Println(err)
	err = errors.New("defer1 error")
}

func e2() {
	var err error
	defer func() {
		fmt.Println(err)
	}()
	err = errors.New("defer2 error")
	return
}
func e3() {
	var err error
	defer func(err error) {
		fmt.Println(err)
	}(err)
	err = errors.New("defer3 error")
	return
}

func main() {
	e1()
	e2()
	e3()
}

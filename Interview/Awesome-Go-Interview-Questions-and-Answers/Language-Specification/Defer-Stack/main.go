package main

import (
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("file")
	defer f.Close()
	if err != nil {
		panic(err.Error())
	}

	b, err := ioutil.ReadAll(f)
	println(string(b))
}
func main2() {
	f, err := os.Open("file")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	println(string(b))
}

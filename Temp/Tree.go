package main

import (
	"bufio"
	"fmt"
)

type Website struct {
	Name string
}

// 定义结构体变量
var site = Website{Name: "studygolang"}

func main() {
	fmt.Printf("%v \n", site)
	fmt.Printf("%+v \n", site)
	fmt.Printf("%#v \n", site)
	fmt.Printf("%d", 10)
	bufio.NewReader().ReadString()
}

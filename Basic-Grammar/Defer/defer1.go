package main

import "fmt"

/**

func f1() (r int) {
	t := 5

	1.直接赋值
	r = t

	2.defer被插入到赋值与返回之间执行，这个例子中返回值没有被修改过
	defer func() {
		t = t + 5
	}()
	3. 空的return指令
	return t
}

*/

func f1() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
func main() {
	fmt.Println(f1())
}

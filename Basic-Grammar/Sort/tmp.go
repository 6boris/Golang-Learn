package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // 未排序的切片数据
	sort.Ints(s)
	fmt.Println(s)
	fmt.Println(sort.SearchInts(s, 5)) //将会输出0而不是1
}

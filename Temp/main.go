package main

import "fmt"

func main() {
	pase_student()
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
		fmt.Printf("%p ", &stu)
		//fmt.Printf("%v ",&stu.Age)
		//fmt.Printf("%v",&stu.Name)
		//fmt.Println(stu)
	}
	v, ok := m["li"]

	fmt.Println(ok, v.Age)
}

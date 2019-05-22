package main

type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func main() {
	s := S{}
	p := &s
	f(s)  //A
	g(&s) //B
	f(p)  //C
	g(p)  //D
}

// In line ABCD, which ones of them have syntax issues?

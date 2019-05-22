package main

type S struct {
	m string
}

func f() *S {
	return __ //A
}

func main() {
	p := __    //B
	print(p.m) //print "foo"
}

//Fill in the blanks in line A and B, to assure the printed output is "foo"

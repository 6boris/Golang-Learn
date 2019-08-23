package main

func myFunction(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	myFunction(66, 77)
}

/*
SET GOOS=linux  && SET GOARCH=amd64 && go tool compile -S -N -l main.go
*/

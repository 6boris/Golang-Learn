package main

func main() {

out_break:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			print(i, ",", j, " ")
			break out_break
		}
		println()
	}
}

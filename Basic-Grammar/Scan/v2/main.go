package main

import "fmt"

func main() {
	var n, sx, sy, ex, ey int
	_, err := fmt.Scanf("%d %d %d %d %d", &n, &sx, &sy, &ex, &ey)
	if err != nil {
		fmt.Println(err)
	}
	ans := max(abs(sx-ex), abs(sy-ey))

	//fmt.Println(n, sx, sy, ex, ey)
	fmt.Println(ans)

}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	number := scanner.Text()
	fmt.Println(number)

	scanner.Scan()
	arr := strings.Split(strings.Trim(scanner.Text(), " "), " ")

	fmt.Println(arr)

	var n, sx, sy, ex, ey int
	_, err := fmt.Scanf("%d %d %d %d %d", &n, &sx, &sy, &ex, &ey)
	if err != nil {
		fmt.Println(err)
	}
}

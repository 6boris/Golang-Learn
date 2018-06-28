package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func BubbleSort(arr []int) []int {
	tmp := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}

func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

func Print(line string) {
	fmt.Println(line)
}

func main() {
	ReadLine("./Sort/data.txt", Print)

	arr := new(Array)
	arr.Got = []int{1, 3, 2, 4, 5}
	arr.Want = []int{1, 2, 3, 4, 5}
	fmt.Println(arr.Got)
	fmt.Println(arr.Want)
}

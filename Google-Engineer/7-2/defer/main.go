package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writerFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	fmt.Println(writer, "demo")

	//for i := 0; i < 20; i++ {
	//	fmt.Println(writer, i)
	//}

}

func main() {
	writerFile("Google-Engineer/7-2/defer/fib.txt")
	//tryDefer()
	//writerFile("fib.txt")
}

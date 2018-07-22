package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

var data = make([]Array, 0, 2000)

type Array struct {
	Got  []int "json:got"
	Want []int "json:want"
}

func processLine(line []byte) {
	//os.Stdout.Write(line)
	//line := string(line[:])
	//fmt.Printf("%s", line)
	//fmt.Println(bytes.Fields(line))
	//fmt.Println(reflect.TypeOf(string(line)))
	str := string(line)
	fmt.Println(strings.Split(str, " "))

}

func ReadLine(filePth string, hookfn func([]byte)) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()

	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		hookfn(line)    //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

func main() {
	//ReadLine("./Sort/data.txt", processLine)

	str1 := "[5,3,2,5,1] [1,2,3,4,5]"
	str2 := strings.Split(str1, " ")
	//fmt.Println(str2[0])
	//fmt.Println(str2[1])
	fmt.Println(reflect.TypeOf(strings.Split(str1, " ")))
	str2[0] = strings.TrimPrefix(str2[0], "[")
	str2[0] = strings.TrimSuffix(str2[0], "]")
	fmt.Println(str2[0])
	str3 := strings.Split(str2[0], " ")

	fmt.Println(str3)
	fmt.Println(reflect.TypeOf(str2[0]))

	//f, err := os.Open("./Sort/data.txt")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer f.Close()
	//bfRd := bufio.NewReader(f)
	//for {
	//	line, err := bfRd.ReadBytes('\n')
	//	//hookfn(line) //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
	//	fmt.Printf("%s", line)
	//	if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
	//		if err == io.EOF {
	//
	//		}
	//	}
	//}
}

package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("defer main")
	var user = ""

	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("recover success. err:", err)
			}
		}()

		func() {
			defer func() {
				fmt.Println("defer here")
			}()

			if user == "" {
				panic("should set uer env")
			}

			fmt.Println("after panic")

		}()

	}()

	time.Sleep(5000)
	fmt.Println("end of main function")

}

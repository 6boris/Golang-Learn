package main

import (
	"errors"
	"fmt"
)

type PayInterface interface {
	Do1()
	Do2()
	Do3()
}

type PayChannel1 struct {
}

func (p *PayChannel1) Do1() {
	fmt.Println("PayChannel1:Do1")
}

func (p *PayChannel1) Do2() {

}

func (p *PayChannel1) Do3() {

}

func GetPayInstance(s string) (PayInterface, error) {
	switch s {
	case "1":
		return new(PayChannel1), nil
	default:
		return nil, errors.New("s is err")
	}
}

func main() {
	ch, err := GetPayInstance("11")
	if err != nil {
		fmt.Println(err.Error())
	}
	if ch != nil {
		ch.Do1()
	}
}

package main

import (
	"fmt"
	"encoding/json"
)

type MSResult struct {
	Code 	uint		`json:"code"`
	//Data 	struct{}	`json:"data"`
	Data 	string		`json:"data"`
	Message	string		`json:"message"`
}

type User struct {
	Id		int8
	Name	string
}
func (self *User) ToJson() string {
	res,_ := json.Marshal(self)
	return string(res)
}

func (self *MSResult) ToJson() string {
	res,_ := json.Marshal(self)
	return string(res)
}

func main() {
	m := new(MSResult)
	m.Code = 200
	//m.Data = new()
	m.Message = "注册成功"
	u := User{1,"KYLE"}
	m.Data = u.ToJson()

	fmt.Println(u.ToJson())
	fmt.Println(m.ToJson())
}
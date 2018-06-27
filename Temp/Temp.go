package main

import "fmt"

type Response struct {
	Code    int
	Result  []byte
	Headers map[string]string
}

func (r Response) GetAttrCode() int {
	return r.Code
}
func (r Response) GetAttrResult() []byte {
	return r.Result
}

func (r Response) GetAttrHeader() map[string]string {
	return r.Headers
}

func (r *Response) SetCode(code int) {
	r.Code = code
}

func (r *Response) SetHeaders(key, value string) {
	r.Headers[key] = value
}

type Requests struct {
	Url    string
	Params string
}

type CollectionRequests struct {
	CollectionNumber int
	Requests
	Response
}

func exampleCollectionRequests() {

	var collectionRequests CollectionRequests
	collectionRequests.CollectionNumber = 10
	collectionRequests.Url = "https://www.example.com"
	collectionRequests.Params = "name"
	collectionRequests.Code = 201
	collectionRequests.Result = []byte("hello Golang")

	var headers map[string]string
	headers = make(map[string]string)
	headers["status"] = "Good"
	collectionRequests.Headers = headers
	fmt.Println(collectionRequests)

	fmt.Println(collectionRequests.GetAttrCode())
}

type OtherRequests struct {
	Request Requests
	Resp    Response
	Code    int
}

func (o OtherRequests) GetAttrCode() {
	fmt.Println(fmt.Sprintf("Outer Code = %d", o.Code))
	fmt.Println(fmt.Sprintf("inner Code = %d", o.Resp.Code))
}
func exampleOtherRequests() {
	var other OtherRequests
	other.Code = 201
	other.Resp.Code = 202
	fmt.Println(other)
	other.GetAttrCode()
	fmt.Println(other.Resp.GetAttrCode())
}
func main() {
	exampleCollectionRequests()
	exampleOtherRequests()
}

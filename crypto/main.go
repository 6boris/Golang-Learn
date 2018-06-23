package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	hash := md5.New()
	b := []byte("jiavv5..")
	hash.Write(b)
	fmt.Printf("%x %x\n", hash.Sum(nil), md5.Sum(b))
	hash.Write(nil)
	fmt.Printf("%x %x\n", hash.Sum(b), hash.Sum(nil))
}

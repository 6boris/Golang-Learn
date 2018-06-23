package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/json-iterator/go"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err = json_iterator.Marshal(group)
	os.Stdout.Write(b)
}

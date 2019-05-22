package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

func main() {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", "package main; var a = 0", parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	var v visitor
	ast.Walk(v, f)
}

type visitor int

func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	var s string
	switch node := n.(type) {
	case *ast.Ident:
		s = node.Name
	case *ast.BasicLit:
		s = node.Value
	}
	fmt.Printf("%s%T: %s\n", strings.Repeat("\t", int(v)), n, s)
	return v + 1
}

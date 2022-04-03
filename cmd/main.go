package main

import (
	"fmt"
	"gohexproject/internal/adapters/core/arithmetic"
)

func main() {
	fmt.Println("Hello World")
	arithAdaptor := arithemtic.NewAdapter()
	arithAdaptor.Addition()
}
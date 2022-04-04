package main

import (
	"fmt"
	"hex/internal/adapters/core/arithmetic"
)

func main() {
	fmt.Println("Hello World")
	arithAdaptor := arithmetic.NewAdapter()
	additionResult, error := arithAdaptor.Addition(2, 4)
	if error != nil {
		fmt.Println(additionResult)
	}
	println(error)
	
	// arithAdaptor.Subtraction()
	// arithAdaptor.Multiplication()
	// arithAdaptor.Division()
}
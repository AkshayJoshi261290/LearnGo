package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome!")
	a := 5
	b := &a
	fmt.Println("Value of A is: ", a)
	fmt.Println("Value of B/Memory Address of A is: ", b)
	fmt.Printf("Type of b(address of A) is: %T\n", b)

	//read value from mem address
	fmt.Println("Value of A with address is: ", *b)
	fmt.Println("Value of A with address is: ", *&a)

	//change value with pointer
	*b = 10
	fmt.Println("Value of A is: ", a)

}

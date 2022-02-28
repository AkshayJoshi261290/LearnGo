package main

import (
	"fmt"

	CalMath "github.com/AkshayJoshi261290/LearnGo/03_Packages_Functions/Calc"
)

var name string = "Akshay"

func main() {
	fmt.Println("Hello!", "\nI am "+name)
	fmt.Println("Enter Value of A:")
	var A, B float64
	fmt.Scanln(&A)
	fmt.Println("Enter Value of B:")
	fmt.Scanln(&B)
	var add, sub, mul, div = CalMath.CalcFunc(A, B)
	fmt.Println("A+B is", add)
	fmt.Println("A-B is", sub)
	fmt.Println("A*B is", mul)
	fmt.Println("A/B is", div)
}

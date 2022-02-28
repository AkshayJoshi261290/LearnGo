package main

import (
	"fmt"
	"os"
	"strings"
)

//FizzBuzz example
func main() {

	var usrInp string = ""
	var maxlim int = 0
StartOfScript:
	fmt.Println("------Start Of Script-------")
	fmt.Println("Hello There!", "\n Do you want to see FizzBuzz(Yes/No)?")
	fmt.Scanln(&usrInp)
	usrInp = strings.Trim(strings.ToUpper(usrInp), " ")
	switch usrInp {
	case "YES":
		fmt.Println("Please enter maximum limit: ")
		fmt.Scanln(&maxlim)
		ShowFizzBuzz(maxlim)
	case "NO":
		os.Exit(1)
	default:
		fmt.Println("Wrong input! please enter Yes or No to preceed further.")
		goto StartOfScript
	}
	fmt.Println("------End Of Script-------")
}
func ShowFizzBuzz(lim int) {
	for i := 1; i <= lim; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

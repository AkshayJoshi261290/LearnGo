package main

import (
	"fmt"
)

//a closure is a function that references variables outside of its scope
func main() {
	fmt.Println("Welcome!")

	//call normal function
	send("This is normal function!")

	//lets create a Anonymous function for normal function
	//called directly by passing parameters ()
	//This function is similar to a regular function — except it does not have a name.
	func(message string) {
		fmt.Println(message)
	}("This is Anonymous function of normal function!")

	//Call give_me_a_func- A function that returns a function
	send_func := give_me_a_func()                                //variable stores the function
	send_func("Calling function that calling annomous function") //call that variable like function

	//A closure is a function that references variables outside its own function body (scope).
	inc := increamentor()
	for i := 0; i < 10; i++ {
		fmt.Println("Incrementor's i:", inc()) //it keeps the old value of i outside the scope
	}
}

func increamentor() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func give_me_a_func() func(string) {
	//A function that returns a function
	return func(message string) {
		fmt.Println(message)
	}
}

func send(message string) {
	// to understand Anonymous function lets create normal function
	fmt.Println(message)
}

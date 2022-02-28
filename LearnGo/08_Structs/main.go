package main

import (
	"fmt"
	"strconv"
)

//decalration of struct
type Person struct {
	Fname  string
	Lname  string
	city   string
	gender string
	age    int
}

//Method to write shortintro for persons(value receiver)
func (p Person) Shortintro() string {
	return "Hello There!, My name is " + p.Fname + " " + p.Lname + " and I am " + strconv.Itoa(p.age) + "."
}

//Method to write hasbirthday for persons(pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

//Method to write GetMarried for persons(pointer receiver)
func (p *Person) GetMarried(spouceLname string) {
	if p.gender == "M" {
		return
	} else {
		p.Lname = spouceLname
	}
}

//Struct with methods of pointer receivers/ value receivers
func main() {
	fmt.Println("Welcome!")
	//Initializing struct
	p1 := Person{Fname: "Akshay", Lname: "Joshi", city: "Mumbai", gender: "M", age: 30}

	//Another way of initilizing struct
	p2 := Person{"Shivani", "Khawase", "Mumbai", "F", 28}

	//Print Original struct
	fmt.Println("Original values of Person1 & Person2 ares: ", p1, p2)

	//ptrint single variable
	fmt.Println("Person1's Full Name is: " + p1.Fname + " " + p1.Lname)

	//calling method shortintro- value receiver
	fmt.Println(p1.Shortintro())
	fmt.Println(p2.Shortintro())

	//calling method hasbirthday- pointer receiver
	p1.hasBirthday()
	fmt.Println(p1.Shortintro())
	p2.hasBirthday()
	fmt.Println(p2.Shortintro())

	//calling method for getmarried
	p2.GetMarried(p1.Lname)
	p1.GetMarried("TestString")
	fmt.Println(p1.Shortintro())
	fmt.Println(p2.Shortintro())
}

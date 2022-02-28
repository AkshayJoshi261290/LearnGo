package main

import (
	"fmt"
	"math"
)

//2 tructs that represents circle & rectangle, with their area calculation methods
type circle struct {
	x, y, radious float64
}

type rectangle struct {
	height, width float64
}

//value rec mehod for circle
func (c circle) area() float64 {
	return math.Pi * c.radious * c.radious
}

//value rec mehod for rectangle
func (r rectangle) area() float64 {
	return r.height * r.width
}

//shape is an interface to different structs which gives respetive area
type shape interface {
	area() float64
}

//function which takes interface & return area
func getArea(s shape) float64 {
	return s.area()
}

func main() {
	fmt.Println("Welcome!")

	//initialize circle & rectangle
	cr := circle{0, 0, 20}
	rect := rectangle{10, 20}

	//print area using getArea method
	fmt.Println("Area of Circle is: ", getArea(cr))
	fmt.Println("Area of Rectangle is: ", getArea(rect))
}

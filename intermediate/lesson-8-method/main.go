package main

import "fmt"

type Shape struct {
	Rectangle
}

type Rectangle struct {
	length float64
	width  float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Metho with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

func main() {
	rect := Rectangle{
		length: 10,
		width:  9,
	}

	area := rect.Area()
	fmt.Println("Area of rectangle with width 9 and length 10 is: ", area)

	rect.Scale(2)
	area = rect.Area()
	fmt.Println("Area of rectangle with width 9 and lenght 10 is: ", area)

	num := Myint(-5)
	num1 := Myint(9)

	fmt.Println(num.isPositive())
	fmt.Println(num1.isPositive())
	fmt.Println(num.welcomeMessage())

	s := Shape{Rectangle: Rectangle{length: 14, width: 12}}
	fmt.Println(s.Area())
	fmt.Println(s.Rectangle.Area())
}

type Myint int

// Method on user-defined type
func (m Myint) isPositive() bool {
	return m > 0
}

func (Myint) welcomeMessage() string {
	return "Welcome to MyInt Type"
}

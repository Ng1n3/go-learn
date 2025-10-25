package main

import "fmt"

func main() {
	// Printing Functions
	fmt.Print("Hello ")
	fmt.Print("World!")
	fmt.Print(12, 456)

	fmt.Println("Hello ")
	fmt.Println("World!")
	fmt.Println(12, 456)

	name := "John"
	age := 25
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Binary: %b, Hex: %X\n", age, age)

	// Formatting Functions
	s := fmt.Sprint("Hello ", "World!", 123, 456)
	fmt.Println(s)

	s = fmt.Sprintln("Hello", "World!", 123, 456)
	fmt.Print(s)
	fmt.Print(s)

	sf := fmt.Sprintf("Name: %s, Age: %d\n", name, age)
	fmt.Println(sf)

	// Scanning Functions
	// var nm string
	// var ag int

	fmt.Print("Enter your name and age: ")
	// fmt.Scan(&nm, &ag)
	// fmt.Scanln(&nm, &ag)
	// fmt.Scanf("%s %d", &nm, &ag)
	// fmt.Printf("Name: %s, Age: %d\n", nm, ag)

	// Error Formatting Functions
	err := checkAge(15)
	if err != nil {
		fmt.Println(err)
	}
}

func checkAge(age int) error {
	if age > 18 {
		return fmt.Errorf("age %d is too young to drive. ", age)
	}
	return nil
}

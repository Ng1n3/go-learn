package main

import "fmt"

func main() {

	sum := add(1, 2)
	fmt.Println(sum)

	func() {
		fmt.Println("Hello anonymous")
	}()

	greet := func() {
		fmt.Println("Hello there and welcome to functions in GO")
	}

	greet()

	operation := add

	result := operation(3, 5)

	fmt.Println(result)

	// passing a function as an arguement
	result2 := applyOperation(5, 3, add)

	fmt.Println("5 + 3", result2)

	// Returning and using a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2 = ", multiplyBy2(6))
}

func add(a, b int) int {
	return a + b
}

// Function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func createMultiplier(factore int) func(int) int {
	return func(x int) int {
		return x * factore
	}
}

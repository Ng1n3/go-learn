package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))

	fmt.Println(sumOfDigits(25))
}

func factorial(n int) int {

	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
	// n *(n -1) * (n -2) * factorial (n - 3)...... factorial(0)
}

func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}

	return n%10 + sumOfDigits(n/10)
}

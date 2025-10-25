package main

import "fmt"

func main() {

	message := "Hello, \nGo!"
	rawMessage := `Hello\nGo`
	message1 := "Hello, \tGo!"
	message2 := "Hello, \rGo!"

	fmt.Println(message)
	fmt.Println(rawMessage)
	fmt.Println(message1)
	fmt.Println(message2)

	fmt.Println("Length of message variable is", len(message))

	fmt.Println("The first character in message var is", message[0]) // ASCII

	// Concatination
	greeting := "Hello"
	name := "Alice"
	fmt.Println(greeting + name)

	// Comparison
	str1 := "Apple"
	str2 := "banana"
	str3 := "app"

	fmt.Println(str1 < str2)
}

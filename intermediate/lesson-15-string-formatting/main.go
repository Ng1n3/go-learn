package main

import "fmt"

func main() {

	num := 42
	fmt.Printf("%05d\n", num)

	message := "Hello"
	fmt.Printf("|%10s|\n", message)
	fmt.Printf("|%-10s|\n", message)

	message1 := "Hello \nWorld!"
	message2 := `Hello \nworld!`

	fmt.Println(message1)
	fmt.Println(message2)

	// Where to use backticks
	sqlQuery := `SELECT * FROM users WHERE age > 30`
	fmt.Println(sqlQuery)

}

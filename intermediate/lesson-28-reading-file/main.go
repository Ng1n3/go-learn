package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	newData, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}

	_, err = newData.WriteString("Welcome to the world of Golang!")
	if err != nil {
		fmt.Println("There was an error writing string", err)
		return
	}

	//open file
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer func() {
		fmt.Println("Closing open file")
		file.Close()
	}()

	fmt.Println("File open successfully.")

	// Read the contents of the opened file.
	data := make([]byte, 1024)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Error reading data from file: ", err)
		return
	}

	fmt.Println("File content: ", string(data))

	// Using Bufio
	scanner := bufio.NewScanner(file)

	// Rad line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line: ", line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
}

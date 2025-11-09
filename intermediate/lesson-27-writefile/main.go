package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("text.txt")
	if err != nil {
		fmt.Println("Error creating file: ", err)
	}
	defer file.Close()

	// write data to file
	data := []byte("Hello world!\n")
	bt, err := file.Write(data)
	if err != nil {
		fmt.Println("error writing to file", err)
		return
	}
	fmt.Println("Bytes created", bt)

	_, err = file.WriteString("welcome to Golang\n")
	if err != nil {
		fmt.Println("Error writing to string")
	}
}

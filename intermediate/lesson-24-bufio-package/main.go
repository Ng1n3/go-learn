package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(strings.NewReader("Hello, bufio packageee!\nHow are you doing?"))

	// Reading byte slice
	data := make([]byte, 20)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading string: ", err)
		return
	}

	fmt.Println("Read string: ", line)

	writer := bufio.NewWriter(os.Stdout)

	// Writing byte slice
	data2 := []byte("Hello, bufio package!\n")
	n2, err := writer.Write(data2)
	if err != nil {
		fmt.Println("Error writing: ", err)
	}
	fmt.Printf("Wrote %d bytes\n", n2)

	// Flush the buffer to ensure all data is written to os.Stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("There was an error flushing writer", err)
		return
	}

	// Writing string
	str := "This is a string.\n"
	n3, err := writer.WriteString(str)
	if err != nil {
		fmt.Println("There was an error writing string", err)
		return
	}

	fmt.Printf("Wrote %d bytes. \n", n3)

	// Flush the buffer
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flusing writer", err)
		return
	}

	fmt.Println("There.")
}

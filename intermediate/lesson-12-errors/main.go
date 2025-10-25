package main

import (
	"errors"
	"fmt"
)

type myError struct {
	message string
}

func (me *myError) Error() string {
	return fmt.Sprintf("error: %s", me.message)
}

func eProcess() error {
	return &myError{"Custom error message"}
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("math: square root of negativ number")
	}
	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("empty data")
	}
	// process data
	return nil
}

func readData() error {
	if err := readConfig(); err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

func readConfig() error {
	return errors.New("config error")
}

func main() {
	result, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	result1, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result1)

	data := []byte{}
	if err := process(data); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Data Processed Successfully")

	if err1 := eProcess(); err1 != nil {
		fmt.Println("Error: ", err1)
		return
	}

	if err := readData(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data read Successfully")

}

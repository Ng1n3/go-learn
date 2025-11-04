package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := []byte("Hello, Base64 Encoding")

	// Encode Base64
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	// Decode from Base64
	bs, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("There was an error decoding", err)
		return
	}

	fmt.Println("Decoded string: ", string(bs))
}

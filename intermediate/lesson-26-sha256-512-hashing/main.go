package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/ascii85"
	"encoding/base64"
	"fmt"
	"io"

	base91 "github.com/Equim-chan/base91-go"
)

func main() {

	password := "Password12345"

	hash := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password))

	fmt.Println(password)
	fmt.Println("SHA256: ", hash)
	fmt.Println("SHA-512: ", hash512)
	fmt.Printf("SHA-256 Hash hex val: %x\n", hash)
	fmt.Printf("SHA-512 Has hex val: %x\n", hash512)

	fmt.Println("#------------USING SALT-----------#")
	salt, err := generateSalt()
	fmt.Println("Original Salt: ", salt)
	fmt.Printf("Original Salt hex val: %x\n: ", salt)
	if err != nil {
		fmt.Println("There was an error generating salt", err)
		return
	}

	// Hash the password with salt
	hashedPassword := hashPassword(password, salt)
	fmt.Println("Hashed password: ", hashedPassword)

	//----- Base 64 Endoding -----
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Base64 Enconding: Salt: ", saltStr)

	//----- Base 85 Encoding -----
	dst := make([]byte, ascii85.MaxEncodedLen(len(hashedPassword)))
	n := ascii85.Encode(dst, hashedPassword)
	fmt.Println("Base85 Encoded", dst[:n])

	// ----- Base 91 Encoding -----
	dst2 := make([]byte, int(float64(len(hashedPassword))*1.25))
	n2 := base91.Encode(dst2, hashedPassword)
	fmt.Println("Base91 Encoded: ", n2)

	// Verify
	// retrieve the saltStr and decode it
	decodeSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Unable to decode salt: ", err)
		return
	}
	loginHash := hashPassword(password, decodeSalt)

	// Compare the stored hashed and loginHash
	// subtle.ConstantTimeCompare(hashedPassword, loginHash) == 1
	if bytes.Equal(hashedPassword, loginHash) {
		fmt.Println("Password is correct. You are logged in.")
	} else {
		fmt.Println("Password incorrectj")
	}
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// Function to Hash password
func hashPassword(password string, salt []byte) []byte {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	// return base64.StdEncoding.Strict().EncodeToString(hash[:])
	return hash[:]

}

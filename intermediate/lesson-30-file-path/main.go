package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	relativePath := "./data/file.txt"
	absolutePath := "/home/user/docs/file.txt"

	joinedPath := filepath.Join("home", "Documents", "downloads", "file.zip")
	fmt.Println("joined path: ", joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Cleaned path: ", normalizedPath)

	dir, file := filepath.Split("/home/user/docs/file.txt")
	fmt.Println("file: ", file)
	fmt.Println("Dir: ", dir)
	fmt.Println(filepath.Base("/home/user/docs"))

	// check for absolute path
	fmt.Println("Is relativePath variable absolute:", filepath.IsAbs(relativePath))
	fmt.Println("Is absolutePath variable absolute: ", filepath.IsAbs(absolutePath))

	// Extension of the path
	fmt.Println(filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		panic(err)
	}

	fmt.Println(rel)

	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("absoute path: ", absPath)
}

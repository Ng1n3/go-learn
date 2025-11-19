package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	//err := os.Mkdir("subdire", 0755)
	//checkError(err)
	//	checkError(os.Mkdir("subdir", 0755))

	//defer os.RemoveAll("subdir")

	//	os.WriteFile("subdir/file", []byte(""), 0755)

	checkError(os.MkdirAll("subdir/parent/child", 0755))

	os.WriteFile("subdir/parent/child/file", []byte(""), 0755)

	result, err := os.ReadDir("subdir/parent")
	checkError(err)

	for _, v := range result {
		fmt.Println(v.Name(), v.IsDir(), v.Type())
	}

	checkError(os.Chdir("subdir/parent/child"))

	result, err = os.ReadDir(".")
	checkError(err)

	for _, v := range result {
		fmt.Println(v.Name(), v.IsDir(), v.Type())
	}

	checkError(os.Chdir("../../.."))
	dir, err := os.Getwd()
	checkError(err)

	fmt.Println(dir)

	// filepath.Walk and filepath.WalkDir
	fmt.Println("-----------Walking directory-----------")
	pathfile := "subdir/parent/child"
	err = filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}
		fmt.Println(path)
		return nil
	})

	checkError(err)

	checkError(os.RemoveAll("./subdir"))
}

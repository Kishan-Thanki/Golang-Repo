package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("sample_file.txt")
	if err != nil {
		// Log the error
		panic(err)
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Is folder:", fileInfo.IsDir())
	fmt.Println("File size:", fileInfo.Size())
	fmt.Println("File permissions:", fileInfo.Mode())
	fmt.Println("File modified at:", fileInfo.ModTime())

	// Read file content
	fmt.Println()

	buf := make([]byte, fileInfo.Size())
	d, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	println("Data size:", d)
	for i := 0; i < len(buf); i++ {
		print(string(buf[i]))
	}

	// Second way to read file content
	// data, err := os.ReadFile("sample_file.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	// To read folders
	fmt.Println("")
	fmt.Println("")
	dir, err := os.Open(".")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	dirInfo, err := dir.ReadDir(2)
	if err != nil {
		panic(err)
	}

	for _, fi := range dirInfo {
		fmt.Println(fi.Name())
	}

	// To create a file
	newFile, err := os.Create("sample_file2.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	newFile.WriteString("Hello, Python")
}

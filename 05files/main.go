package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("do sommething with files")

	content := "there may be something info abvailablee inf d mbjbhjgdisdgsjdhs"
	file, err := os.Create("./examplefile.txt")

	checkNilError(err)
	length, err := io.WriteString(file, content)

	checkNilError(err)
	fmt.Println("length of file", length)

	defer file.Close()
	readFile("./examplefile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilError(err)
	fmt.Println("raw data in file ", databyte)
	fmt.Println("data in file ", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

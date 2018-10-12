package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := getCmdLineArguments(1)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Could not open file \"%v\"!", fileName)
		os.Exit(-1)
	}

	io.Copy(os.Stdout, file)

	file.Close()
}

func getCmdLineArguments(pos int) string {
	if !(pos > len(os.Args)-1) {
		return os.Args[pos]
	}
	return ""
}

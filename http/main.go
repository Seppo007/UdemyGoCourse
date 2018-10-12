package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Printf("There was an error: %v", err)
		os.Exit(-1);
	}

	fmt.Printf("Response status is: %s\n", resp.Status)

	io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
}

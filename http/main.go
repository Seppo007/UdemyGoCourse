package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.de")
	if err != nil {
		fmt.Printf("There was an error: %v", err)
		os.Exit(-1);
	}
	p := make([]byte, 12000);
	fmt.Printf("Response status is: %s\n", resp.Status)
	resp.Body.Read(p)
	fmt.Printf("Body is: %s\n", p)
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Printf("There was an error: %v", err)
		os.Exit(-1);
	}

	fmt.Printf("Response status is: %s\n", resp.Status)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
	resp.Body.Close()
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Just wrote %v bytes\n", len(bs))
	return len(bs), nil
}

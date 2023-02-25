package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// makes a request
	// response saved to res
	// or error saved to err
	res, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// make a byte slice with n number of empty elements in side of it
	bs := make([]byte, 99999)
	// takes body and puts inside byte slice
	res.Body.Read(bs)
	// this displays response body as a string (response body is the HTML of the Google home page)
	fmt.Println(string(bs))
}

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

	fmt.Println(res)
}

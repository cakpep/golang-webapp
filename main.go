package main

import (
	"fmt"
	"net/http"
	"webapp/src"
)

func main() {

	// Run Routes
	src.Routes()

	// After defining our server, we finally "listen and serve" on port 3000
	// The second argument is the handler, which we will come to later on, but for now it is left as nil,
	// and the handler defined above (in "HandleFunc") is used
	fmt.Println("server started at localhost:3000")
	http.ListenAndServe(":3000", nil)
}

package src

import (
	"fmt"
	"net/http"
	controller "webapp/src/controllers"
)

func Routes() {
	// The "HandleFunc" method accepts a path and a function as arguments
	// (Yes, we can pass functions as arguments, and even trat them like variables in Go)
	// However, the handler function has to have the appropriate signature (as described by the "handler" function below)
	http.HandleFunc("/", controller.ActionIndex)
	http.HandleFunc("/about", controller.ActionAbout)

	//http.Handle("/", http.FileServer(http.Dir("src/assets")))

	// Handle the assets
	http.Handle("/assets/",
		http.StripPrefix("/assets/",
			http.FileServer(http.Dir("src/assets"))))

}

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
func PageIndexHandler2(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World!")
}

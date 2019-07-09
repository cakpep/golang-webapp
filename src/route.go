package src

import (
	"fmt"
	"net/http"
	ctrl "webapp/src/controllers"
)

func Routes() {
	// The "HandleFunc" method accepts a path and a function as arguments
	// (Yes, we can pass functions as arguments, and even trat them like variables in Go)
	// However, the handler function has to have the appropriate signature (as described by the "handler" function below)

	// Site Controller
	http.HandleFunc("/", ctrl.SiteIndex)
	http.HandleFunc("/about", ctrl.SiteAbout)


	// Person Controller
	http.HandleFunc("/persons", ctrl.PersonIndex)

	// Handle the assets
	http.Handle("/assets/",
		http.StripPrefix("/assets/",
			http.FileServer(http.Dir("src/assets"))))

}

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
func PrintText(text string, w http.ResponseWriter, r *http.Request) {
	//Print text variable into browser
	fmt.Fprintf(w, text)
}

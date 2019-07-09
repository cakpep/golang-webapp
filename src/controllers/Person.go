package controllers

import (
	"net/http"
	model "webapp/src/models"
)

/**
	Person Index
*/
func PersonIndex(w http.ResponseWriter, r *http.Request) {
	var data = model.Person{
		Name:    "Bruce Wayne",
		Gender:  "male",
		Hobbies: []string{"Reading Books", "Traveling", "Buying things"},
		Info:    model.Info{"Wayne Enterprises", "Gotham City"},
	}

	ViewHandler("person",data,w,r)
}

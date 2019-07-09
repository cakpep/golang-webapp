package controllers

import (
	"net/http"
)

/**
	Action Index
 */
func ActionIndex(w http.ResponseWriter, r *http.Request) {
	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Batman",
	}
	ViewHandler("about",data,w,r)
}

/**
Action Index
*/
func ActionAbout(w http.ResponseWriter, r *http.Request) {
	var data = map[string]interface{}{
		"title": "About Learning Golang Web",
		"name":  "Batman",
	}
	ViewHandler("about",data,w,r)
}
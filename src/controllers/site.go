package controllers

import (
	"net/http"
)

/**
	Action Index
 */
func SiteIndex(w http.ResponseWriter, r *http.Request) {
	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Batman",
	}
	ViewHandler("index",data,w,r)
}

/**
Action Index
*/
func SiteAbout(w http.ResponseWriter, r *http.Request) {
	var data = map[string]interface{}{
		"title": "About Learning Golang Web",
		"name":  "Batman",
	}
	ViewHandler("about",data,w,r)
}
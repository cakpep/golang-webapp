package controllers

import (
	"text/template"
	"net/http"
)

func ViewHandler(fileName string, data interface{}, w http.ResponseWriter, r *http.Request) {

	err := template.Must(template.ParseGlob("src/views/*/*")).ExecuteTemplate(w, fileName, data)
	//err := template.Must(template.ParseFiles(
	//	//"src/views/site/"+ fileName +".html",
	//	"src/views/layouts/_header.html",
	//	"src/views/layouts/_footer.html",
	//)).ExecuteTemplate(w, fileName, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err.Error())
		return
	}

}

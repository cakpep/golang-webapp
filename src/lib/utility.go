package lib

import (
	"html/template"
	"net/http"
)

func ViewHandler(fileName string, data interface{}, w http.ResponseWriter, r *http.Request) {

	var tmpl = template.Must(template.ParseFiles(
		//"src/views/site/"+ fileName +".html",
		"src/views/layouts/_header.html",
		"src/views/layouts/_footer.html",
	))

	var err = tmpl.ExecuteTemplate(w, fileName, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
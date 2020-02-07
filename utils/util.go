package utils

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"workplace/IMGo/structManager"
)

func Resp(writer http.ResponseWriter, responseFormat structManager.H) {

	res, err := json.Marshal(responseFormat)
	if err != nil {
		log.Println(err.Error())
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

// html page render
func RegisterView() {
	// handle static files
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	// ** is dir, * is file
	tpl := template.Must(template.ParseGlob("view/**/*"))

	for _, v := range tpl.Templates() {
		tplName := v.Name()
		http.HandleFunc(tplName, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tplName, nil)
		})
	}
}
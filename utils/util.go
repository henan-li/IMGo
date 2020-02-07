package utils

import (
	"fmt"
	//"encoding/json"
	"html/template"
	"log"

	//"log"
	"net/http"
	//"workplace/IMGo/structManager"
)

//func Resp(writer http.ResponseWriter, responseFormat structManager.H) {
//
//	res, err := json.Marshal(responseFormat)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	writer.Header().Set("Content-Type", "application/json")
//	writer.WriteHeader(http.StatusOK)
//	writer.Write(res)
//}

// html page render
func RegisterView() {
	// handle static files
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	// parse other html files under view folder
	tpl, err := template.ParseGlob("view/**/*")
	if nil != err {
		log.Fatal(err)
	}

	for _, v := range tpl.Templates() {
		tplname := v.Name()
		fmt.Println("HandleFunc     " + v.Name())
		http.HandleFunc(tplname, func(w http.ResponseWriter, request *http.Request) {
			fmt.Println("parse     " + v.Name() + "==" + tplname)
			err := tpl.ExecuteTemplate(w, tplname, nil)
			if err != nil {
				log.Fatal(err.Error())
			}
		})
	}
}

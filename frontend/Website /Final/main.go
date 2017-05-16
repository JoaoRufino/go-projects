package main

import (
	"fmt"
	"html/template"
	"net/http"

	"./models"
)

var denm map[string]*models.DENM

// vai buscar o template index e apresenta-o
func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"views/index.html",
		"views/header.html",
		"views/footer.html",
	)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", denm)
}

func main() {
	denm = make(map[string]*models.DENM, 0)
	fmt.Println("Listening on port :80...")
	http.Handle("/components/", http.StripPrefix("/components/", http.FileServer(http.Dir("./components/"))))
	http.HandleFunc("/", index)
	http.ListenAndServe(":80", nil)
}

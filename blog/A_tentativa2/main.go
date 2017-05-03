package main

import (
	"fmt"
	"html/template"
	"net/http"

	"./helpers"
	"./models"
)

var posts map[string]*models.Post

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
	t.ExecuteTemplate(w, "index", posts)
}

func write(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"views/write.html",
		"views/header.html",
		"views/footer.html",
	)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	t.ExecuteTemplate(w, "write", nil)
}

func edit(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"views/write.html",
		"views/header.html",
		"views/footer.html",
	)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	id := r.FormValue("id")
	post, found := posts[id]

	if !found {
		http.NotFound(w, r)
	}
	t.ExecuteTemplate(w, "write", post)
}

func savePost(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	var post *models.Post
	if id != "" {
		post = posts[id]
		post.Title = title
		post.Content = content
	} else {
		post := models.NewPost(helpers.GenerateId(), title, content)
		posts[post.Id] = post
	}
	http.Redirect(w, r, "/", 302)

}

func deletePost(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.NotFound(w, r)
	}
	delete(posts, id)
	http.Redirect(w, r, "/", 302)

}

func main() {
	posts = make(map[string]*models.Post, 0)
	fmt.Println("Listening on port :80...")
	http.Handle("/bower_components/", http.StripPrefix("/bower_components/", http.FileServer(http.Dir("./bower_components/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/write", write)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/savePost", savePost)
	http.HandleFunc("/deletePost", deletePost)
	http.ListenAndServe(":80", nil)
}

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

const (
	PORT = 8080
)

func root(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func login(w http.ResponseWriter, r *http.Request) {

	tpl := template.Must(template.ParseGlob("templates/*.html"))

	switch r.Method {
	case "GET":
		tpl.ExecuteTemplate(w, "layout", nil)
	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Fatal(err.Error())
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		fmt.Printf("[%s]:(%s)\n", username, password)
		fmt.Fprintf(w, "[%s]:(%s)\n", username, password)
	}
}

func main() {
	port := strconv.Itoa(PORT)

	// Routes
	http.HandleFunc("/", root)
	http.HandleFunc("/login", login)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

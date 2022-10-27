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

type WireFrames struct {
	Name string
	Url  string
}

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

func wireframe(w http.ResponseWriter, r *http.Request) {
	wfs := []WireFrames{
		{Name: "Figma (Andres 2)", Url: "https://www.figma.com/file/mRxwNoM82pqq5n35hUcv6Z/Astia-Smish(WF)?node-id=1%3A39"},
		{Name: "Tasin 1", Url: "https://tan7-njit-astial-final-version-app-cdvjua.streamlitapp.com"},
		{Name: "Andres 1", Url: "https://tess-wqing-cs491-astia-streamlit-app-jxnwkj.streamlitapp.com"},
		{Name: "Tasin 2", Url: "https://tan7-njit-streamlit-astia-main-3nkj0j.streamlitapp.com"},
		{Name: "Tasin 3", Url: "https://web.njit.edu/~tan7/campaign.html"},
	}

	tpl := template.Must(template.ParseGlob("templates/*.html"))
	tpl.ExecuteTemplate(w, "head", nil)
	tpl.ExecuteTemplate(w, "wireframes", wfs)
	tpl.ExecuteTemplate(w, "close", nil)
}

// func portal(w http.ResponseWriter, r *http.Request) {
// 	tpl := template.Must(template.ParseGlob("templates/*.html"))
// 	tpl.ExecuteTemplate(w, "head", nil)
// 	tpl.ExecuteTemplate(w, "close", nil)
// }

func attack(w http.ResponseWriter, r *http.Request) {

}

func targets(w http.ResponseWriter, r *http.Request) {

}

func campaign(w http.ResponseWriter, r *http.Request) {

}

func main() {
	port := strconv.Itoa(PORT)

	// Routes
	http.HandleFunc("/", root)
	http.HandleFunc("/login", login)
	http.HandleFunc("/wf", wireframe)
	// http.HandleFunc("/portal", portal)
	http.HandleFunc("/attack", attack)
	http.HandleFunc("/target", targets)
	http.HandleFunc("/campaign", campaign)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	util "github.com/AOrps/cs491-astia/prod/util"
)

const (
	PORT = 8080
)

// Name-Url: Map Name and Url
type NURL struct {
	Name string
	Url  string
}

// func basePage(w http.ResponseWriter, r *http.Request) func() {
// 	return func() {
// 		fmt.Fprintf(w, "yes")
// 	}
// }

func root(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func login(w http.ResponseWriter, r *http.Request) {

	tpl := template.Must(template.ParseGlob("templates/*.html"))

	switch r.Method {
	case "GET":
		tpl.ExecuteTemplate(w, "login", nil)
	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Fatal(err.Error())
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		// fmt.Printf("[%s]:(%s)\n", username, password)
		// fmt.Fprintf(w, "[%s]:(%s)\n", username, password)

		if username == "demo" && password == "demo" {
			http.Redirect(w, r, "/attack", http.StatusSeeOther)
		}
	}
}

func wireframe(w http.ResponseWriter, r *http.Request) {
	wfs := []NURL{
		{Name: "Figma (Andres 2)", Url: "https://www.figma.com/file/mRxwNoM82pqq5n35hUcv6Z/Astia-Smish(WF)?node-id=1%3A39"},
		{Name: "Tasin 1", Url: "https://tan7-njit-astial-final-version-app-cdvjua.streamlitapp.com"},
		{Name: "Andres 1", Url: "https://tess-wqing-cs491-astia-streamlit-app-jxnwkj.streamlitapp.com"},
		{Name: "Tasin 2", Url: "https://tan7-njit-streamlit-astia-main-3nkj0j.streamlitapp.com"},
		{Name: "Tasin 3", Url: "https://web.njit.edu/~tan7/campaign.html"},
	}

	tpl := template.Must(template.ParseGlob("templates/*.html"))
	navs := []NURL{
		{Name: "attack ⚡️", Url: "attack"},
		{Name: "target 📑", Url: "target"},
		{Name: "campaign 📍", Url: "campaign"},
	}

	tpl.ExecuteTemplate(w, "head", nil)
	tpl.ExecuteTemplate(w, "nav", navs)
	tpl.ExecuteTemplate(w, "wireframes", wfs)
	tpl.ExecuteTemplate(w, "close", nil)
}

func attack(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))
	navs := []NURL{
		{Name: "attack ⚡️", Url: "attack"},
		{Name: "target 📑", Url: "target"},
		{Name: "campaign 📍", Url: "campaign"},
	}

	attackTemplates := make(map[string]string)

	attackTemplates["Double Tap"] = "You have been selected to win: Click here"
	attackTemplates["Sick Family"] = "You have someone sick in the family. "
	attackTemplates["Sign in"] = "A computer from a suspicious IP has signed into your gmail account, click here if not you"
	attackTemplates["You Won"] = "You won an aware. Click here if you want to see the terms and conditions."
	attackTemplates["Flu Season"] = "Be sure to take a flu shot. Go to our app or the 'url' to learn about a place nearby"

	switch r.Method {
	case "POST":
		tpl.ExecuteTemplate(w, "head", nil)
		tpl.ExecuteTemplate(w, "close", nil)
	default:
		tpl.ExecuteTemplate(w, "head", nil)
		tpl.ExecuteTemplate(w, "nav", navs)
		tpl.ExecuteTemplate(w, "attack", attackTemplates)
		tpl.ExecuteTemplate(w, "close", nil)
	}

}

func targets(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))
	navs := []NURL{
		{Name: "attack ⚡️", Url: "attack"},
		{Name: "target 📑", Url: "target"},
		{Name: "campaign 📍", Url: "campaign"},
	}

	users := []util.FakeUser


	switch r.Method {
	case "POST":
		tpl.ExecuteTemplate(w, "head", nil)
		tpl.ExecuteTemplate(w, "close", nil)
	default:
		tpl.ExecuteTemplate(w, "head", nil)
		tpl.ExecuteTemplate(w, "nav", navs)
		tpl.ExecuteTemplate(w, "targets", nil)
		tpl.ExecuteTemplate(w, "close", nil)
	}
}

func campaign(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))
	navs := []NURL{
		{Name: "attack ⚡️", Url: "attack"},
		{Name: "target 📑", Url: "target"},
		{Name: "campaign 📍", Url: "campaign"},
	}

	switch r.Method {
	case "POST":
		tpl.ExecuteTemplate(w, "head", nil)
		tpl.ExecuteTemplate(w, "close", nil)
	default:
		tpl.ExecuteTemplate(w, "head", nil)
		tpl.ExecuteTemplate(w, "nav", navs)
		tpl.ExecuteTemplate(w, "campaign", nil)
		tpl.ExecuteTemplate(w, "close", nil)
	}
}

func main() {
	port := strconv.Itoa(PORT)

	// Routes
	http.HandleFunc("/", root)
	http.HandleFunc("/login", login)
	http.HandleFunc("/wf", wireframe)
	http.HandleFunc("/attack", attack)
	http.HandleFunc("/target", targets)
	http.HandleFunc("/campaign", campaign)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Printf("Started web server at:  http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

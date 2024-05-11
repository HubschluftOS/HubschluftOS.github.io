package cmd

import (
	"log"
	"net/http"
	"text/template"
)

const (
	local = "http://localhost"
	port  = ":8000"

	Bold   = "\033[1m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

type Package struct {
	Name        string
	Version     string
	Description string
	Json        string
}

func Html(name string, html string) {
	log.Print(Yellow + "Loading: " + Reset + html + ".html...")
	http.HandleFunc("/"+name, func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./www/" + html + ".html"))

		packages := map[string][]Package{
			"Packages": {
				{Name: "neofetch", Version: "v7.1.0", Description: "A command-line system information tool written in bash 3.2+", Json: "/db/packages/neofetch.json"},
				{Name: "ey", Version: "v7.1.0", Description: "A command-line system information tool written in bash 3.2+", Json: "/d/packages/neofetch.json"},
			},
		}

		tmpl.Execute(w, packages)
	})
}

func Rss(name string, rss string) {
	log.Print(Yellow + "Loading: " + Reset + rss + ".rss")
	http.HandleFunc("/"+name, func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./www/" + rss + ".xml"))
		tmpl.Execute(w, nil)
	})
}

func Server() {
	log.Print(Yellow+"Server started: "+Reset, local, port)
	Html("", "index")
	Html("packages", "pkg")
	Html("download", "download")
	Html("team", "team")
	Rss("rss", "news")

	log.Print(Yellow + "Loading: " + Reset + "/static/...")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./www/static"))))
	log.Print(Yellow + "Loading: " + Reset + "/db/...")
	http.Handle("/db/", http.StripPrefix("/db/", http.FileServer(http.Dir("./db/"))))
	log.Print(Yellow + "Server listening on: " + Reset + local + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

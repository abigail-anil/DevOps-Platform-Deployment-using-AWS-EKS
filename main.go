package main

import (
	"log"
	"net/http"
)

func servePage(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/"+filename)
	}
}

func main() {
	http.HandleFunc("/", servePage("dashboard.html"))
	http.HandleFunc("/dashboard", servePage("dashboard.html"))
	http.HandleFunc("/deployments", servePage("deployments.html"))
	http.HandleFunc("/monitoring", servePage("monitoring.html"))
	http.HandleFunc("/runbooks", servePage("runbooks.html"))
	http.HandleFunc("/about", servePage("about.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
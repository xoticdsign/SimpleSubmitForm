package main

import (
	"fmt"
	"log"
	"net/http"
)

func MainPageHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Main Page -> Serving %q file\n", "index.html")

	http.ServeFile(w, r, "index.html")
}

func SubmitFormHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")

		fmt.Printf("Submit Form -> %v just wrote his %v with a following message: %v\n", name, email, message)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func main() {
	static := http.FileServer(http.Dir("./static"))
	http.Handle("/static", http.StripPrefix("/static", static))

	http.HandleFunc("/", MainPageHandle)
	http.HandleFunc("/submit", SubmitFormHandle)

	fmt.Printf("Server -> Starting a server at :8080\n")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal()
	}
}

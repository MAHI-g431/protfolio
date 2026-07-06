package main

import (
	"fmt"
	"net/http"
)

func Portfolio(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		name := r.FormValue("name")
		email := r.FormValue("email")

		fmt.Println("New Form Submission")
		fmt.Println("Name:", name)
		fmt.Println("Email:", email)

		http.Redirect(w, r, "/?success=true#Contact", http.StatusSeeOther)
		return
	}

	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", Portfolio)

	fmt.Println("Server running at http://localhost:3000")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

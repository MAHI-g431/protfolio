package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", loginPage)
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Server running at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		username := r.FormValue("username")
		password := r.FormValue("password")

		fmt.Println("Username:", username)
		fmt.Println("Password:", password)

		fmt.Fprintln(w, "Login successful!")
	}
}

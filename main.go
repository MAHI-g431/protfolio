package main

import (
	"fmt"
	"net/http"
)

func Protfolio(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		name := r.FormValue("name")
		email := r.FormValue("email")

		fmt.Println("New form submission")
		fmt.Println("Name:", name)
		fmt.Println("Email:", email)
		fmt.Println("_______________________________________________________________________________________")
		http.Redirect(w, r, "/sucess=true#contact", http.StatusSeeOther)
		return
	}
	http.ServeFile(w, r, "index.html")

}
func main() {
	http.HandleFunc("/", Protfolio)
	http.Handle("/tick.png", http.FileServer(http.Dir(".")))
	fmt.Println("Server running at http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}

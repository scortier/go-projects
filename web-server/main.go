package main

import (
	"fmt"
	"log"
	"net/http"
)

// r -> user send request to the server, r is pointing to the request
// w -> server send response to the user
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello from Form	Handler")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	Fname := r.FormValue("fname")
	Lname := r.FormValue("lname")
	fmt.Fprintf(w, "First Name = %s\n", Fname)
	fmt.Fprintf(w, "Last Name = %s\n", Lname)

}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login")
}

func main() {
	fmt.Println("Hello World")
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

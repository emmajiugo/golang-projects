package main

import (
    "fmt"
	"log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello, World!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Received: Name=%s, Address=%s\n", name, address)
	} else {
		http.ServeFile(w, r, "./static/form.html")
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	log.Println("Starting server on :8088")
	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
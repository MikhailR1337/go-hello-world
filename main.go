package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":3000"

func main() {
	http.HandleFunc("/", helloWorldHandler)

	fmt.Println("Hello from server on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not correct", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "./static")
}

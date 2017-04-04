package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloRequest)
	http.HandleFunc("/", getRequest)

	log.Println("Listening on http://localhost:9999/")
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func helloRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Print(w, "<h1>Hello</h1>")
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	file_resquested := "./" + r.URL.Path
	http.ServeFile(w, r, file_resquested)
	return
}

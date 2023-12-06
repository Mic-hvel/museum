package main

import (
	"fmt"
	"net/http"
)

func handleHello (w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello from a Go program"))
}

func main () {
	server := http.NewServeMux()
	http.HandleFunc("/hello", handleHello)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":3333", server)
	if err != nil {
		fmt.Println("Error while opening server")
	}
}
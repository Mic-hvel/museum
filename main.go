package main

import (
	"fmt"
	"net/http"
)

func handleHello (w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello from a Go program"))
}

func main () {
	http.HandleFunc("/hello", handleHello)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Error while opening server")
	}
}
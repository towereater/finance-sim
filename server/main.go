package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Scemo chi legge!")
	})
	http.ListenAndServe(":8080", nil)
}

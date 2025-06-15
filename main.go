package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Chocolate bar Backend")
	})
	fmt.Println("Server running on http://localhost:3000")
	fmt.Println("Server is running and fmt let me print messages to the terminal")
	http.ListenAndServe(":3000", nil)
}

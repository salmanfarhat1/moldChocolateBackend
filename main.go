package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/salmanfarhat1/moldChocolateBackend/api"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", api.HelloHandler).Methods("GET")
	fmt.Println("Server on http://localhost:3000")
	http.ListenAndServe(":3000", r)
}

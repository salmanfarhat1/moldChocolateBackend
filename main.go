package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/salmanfarhat1/moldChocolateBackend/api"
	"github.com/salmanfarhat1/moldChocolateBackend/db"
)

func main() {
	dbConn := db.InitDB()
	defer dbConn.Close()

	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", api.HelloHandler).Methods("GET")
	r.HandleFunc("/chocolates", api.GetChocolatesHandler(dbConn)).Methods("GET") // ✅ Add this line
	r.HandleFunc("/chocolates", api.CreateChocolateHandler(dbConn)).Methods("POST")

	r.HandleFunc("/variants", api.GetVariantsHandler(dbConn)).Methods("GET") // ✅ Add this line
	r.HandleFunc("/chocolates/{id}/variants", api.GetVariantsByChocolateIDHandler(dbConn)).Methods("GET")
	r.HandleFunc("/variants", api.CreateVariantHandler(dbConn)).Methods("POST")

	r.PathPrefix("/photos/").Handler(http.StripPrefix("/photos/", http.FileServer(http.Dir("./photos"))))

	r.PathPrefix("/admin/").Handler(http.StripPrefix("/admin/", http.FileServer(http.Dir("./admin"))))

	fmt.Println("Server on http://localhost:3000")
	http.ListenAndServe(":3000", r)
}

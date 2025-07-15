package api

import (
	"encoding/json"
	"net/http"

	"database/sql"

	"github.com/salmanfarhat1/moldChocolateBackend/db"
)

func GetChocolatesHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		chocolates, err := db.GetChocolates(dbConn)
		if err != nil {
			http.Error(w, "Failed to fetch chocolates", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(chocolates)
	}
}

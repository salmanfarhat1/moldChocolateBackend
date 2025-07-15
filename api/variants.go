package api

import (
	"encoding/json"
	"net/http"

	"database/sql"

	"github.com/salmanfarhat1/moldChocolateBackend/db"
)

func GetVariantsHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		variants, err := db.GetVariants(dbConn)

		if err != nil {
			http.Error(w, "error in retreiving data of the all variants", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(variants)

	}
}

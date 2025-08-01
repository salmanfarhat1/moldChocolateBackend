package api

import (
	"encoding/json"
	"net/http"

	"database/sql"

	"github.com/salmanfarhat1/moldChocolateBackend/db"
	"github.com/salmanfarhat1/moldChocolateBackend/models"
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

func CreateVariantHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var variant models.Variants
		if err := json.NewDecoder(r.Body).Decode(&variant); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		err := db.InsertVariant(dbConn, &variant)
		if err != nil {
			http.Error(w, "Failed to insert variant", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"data":    variant,
		})
	}
}

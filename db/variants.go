package db

import (
	"database/sql"
	"log"

	"github.com/salmanfarhat1/moldChocolateBackend/models"
)

func GetVariants(db *sql.DB) ([]models.Variants, error) {
	rows, err := db.Query("SELECT id, chocolate_id, size, weight, price FROM chocolate_variants")
	if err != nil {
		log.Println("❌ Query failed:", err)
		return nil, err
	}
	defer rows.Close()

	var variants []models.Variants
	for rows.Next() {
		var v models.Variants
		if err := rows.Scan(
			&v.ID,
			&v.ChocolateID,
			&v.Size,
			&v.Weight,
			&v.Price,
		); err != nil {
			log.Println("❌ Scan error:", err)
			return nil, err
		}
		variants = append(variants, v)
	}
	return variants, nil
}

package db

import (
	"database/sql"
	"log"

	"github.com/salmanfarhat1/moldChocolateBackend/models"
)

func GetChocolates(db *sql.DB) ([]models.Chocolate, error) {
	rows, err := db.Query("SELECT id, name, ingredients FROM chocolates")
	if err != nil {
		log.Println("❌ Query failed:", err)
		return nil, err
	}
	defer rows.Close()

	var chocolates []models.Chocolate

	for rows.Next() {
		var choco models.Chocolate
		if err := rows.Scan(&choco.ID, &choco.Name, &choco.Ingredients); err != nil {
			log.Println("❌ Scan failed:", err)
			return nil, err
		}
		chocolates = append(chocolates, choco)
	}
	return chocolates, nil
}

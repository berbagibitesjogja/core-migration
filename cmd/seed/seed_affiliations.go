package main

import (
	"database/sql"
	"math/rand"
	"github.com/brianvoe/gofakeit/v7"
)

func SeedAffiliations(tx *sql.Tx, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO affiliations (name, variant) VALUES ($1, $2) RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	variants := []string{"student", "society", "foundation"}

	for i := 0; i < count; i++ {
		var id int
		err = stmt.QueryRow(
			gofakeit.Company(),
			variants[rand.Intn(len(variants))],
		).Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

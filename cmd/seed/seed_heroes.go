package main

import (
	"database/sql"
	"math/rand"
	"github.com/brianvoe/gofakeit/v7"
)

func SeedHeroes(tx *sql.Tx, beneficiaryIDs []int, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO heroes (beneficiary_id, name, phone) VALUES ($1, $2, $3) RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for i := 0; i < count; i++ {
		var id int
		err = stmt.QueryRow(
			beneficiaryIDs[rand.Intn(len(beneficiaryIDs))],
			gofakeit.Name(),
			gofakeit.Phone(),
		).Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

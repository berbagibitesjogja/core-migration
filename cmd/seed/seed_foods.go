package main

import (
	"database/sql"
	"math/rand"
	"github.com/brianvoe/gofakeit/v7"
)

func SeedFoods(tx *sql.Tx, donationIDs []int, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO foods (donation_id, name, quantity, weight, unit, notes, expired) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	units := []string{"gr", "ml"}

	for i := 0; i < count; i++ {
		var id int
		err = stmt.QueryRow(
			donationIDs[rand.Intn(len(donationIDs))],
			gofakeit.Word(),
			gofakeit.Number(1, 100),
			gofakeit.Number(1, 1000),
			units[rand.Intn(len(units))],
			gofakeit.Sentence(5),
			gofakeit.Bool(),
		).Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

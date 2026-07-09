package main

import (
	"database/sql"
	"math/rand"
	"github.com/brianvoe/gofakeit/v7"
)

func SeedHeroActions(tx *sql.Tx, heroIDs []int, donationIDs []int, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO hero_actions (hero_id, donation_id, status, code, quantity, weight) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	statuses := []string{"done", "ongoing"}

	for i := 0; i < count; i++ {
		var id int
		err = stmt.QueryRow(
			heroIDs[rand.Intn(len(heroIDs))],
			donationIDs[rand.Intn(len(donationIDs))],
			statuses[rand.Intn(len(statuses))],
			gofakeit.UUID(),
			gofakeit.Number(1, 100),
			gofakeit.Number(1, 1000),
		).Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

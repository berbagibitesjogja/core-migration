package main

import (
	"database/sql"
	"math/rand"
	"github.com/brianvoe/gofakeit/v7"
)

func SeedJobSchedules(tx *sql.Tx, sponsorIDs []int, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO job_schedules (sponsor_id, code, receiver, date) VALUES ($1, $2, $3, $4) RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for i := 0; i < count; i++ {
		var id int
		err = stmt.QueryRow(
			sponsorIDs[rand.Intn(len(sponsorIDs))],
			gofakeit.UUID(),
			gofakeit.Name(),
			gofakeit.Date(),
		).Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

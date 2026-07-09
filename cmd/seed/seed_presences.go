package main

import (
	"database/sql"
	"math/rand"
	"github.com/brianvoe/gofakeit/v7"
)

func SeedPresences(tx *sql.Tx, jobSlotIDs []int, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO presences (job_slot_id, title, description, latitude, longitude, max_distance, code, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	statuses := []string{"active", "end"}

	for i := 0; i < count; i++ {
		var id int
		err = stmt.QueryRow(
			jobSlotIDs[rand.Intn(len(jobSlotIDs))],
			gofakeit.JobTitle(),
			gofakeit.Sentence(5),
			gofakeit.Latitude(),
			gofakeit.Longitude(),
			gofakeit.Number(10, 1000),
			gofakeit.UUID(),
			statuses[rand.Intn(len(statuses))],
		).Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

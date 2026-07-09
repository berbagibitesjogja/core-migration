package main

import (
	"database/sql"
	"math/rand"
	"github.com/brianvoe/gofakeit/v7"
)

func SeedJobSlots(tx *sql.Tx, jobScheduleIDs []int, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO job_slots (job_schedule_id, code, name, need, place_and_time) VALUES ($1, $2, $3, $4, $5) RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for i := 0; i < count; i++ {
		var id int
		err = stmt.QueryRow(
			jobScheduleIDs[rand.Intn(len(jobScheduleIDs))],
			gofakeit.UUID(),
			gofakeit.JobTitle(),
			gofakeit.Number(1, 10),
			gofakeit.Address().Address,
		).Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

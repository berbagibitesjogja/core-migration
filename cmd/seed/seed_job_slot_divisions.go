package main

import (
	"database/sql"
	"math/rand"
)

func SeedJobSlotDivisions(tx *sql.Tx, jobSlotIDs []int, divisionIDs []int, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO job_slot_divisions (job_slot_id, division_id) VALUES ($1, $2) ON CONFLICT DO NOTHING RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for i := 0; i < count; i++ {
		var id int
		err = stmt.QueryRow(
			jobSlotIDs[rand.Intn(len(jobSlotIDs))],
			divisionIDs[rand.Intn(len(divisionIDs))],
		).Scan(&id)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

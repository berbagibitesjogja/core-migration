package main

import (
	"database/sql"
	"math/rand"
)

func SeedJobSlotApplicants(tx *sql.Tx, jobSlotIDs []int, userIDs []int, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO job_slot_applicants (job_slot_id, user_id) VALUES ($1, $2) ON CONFLICT DO NOTHING RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for i := 0; i < count; i++ {
		var id int
		err = stmt.QueryRow(
			jobSlotIDs[rand.Intn(len(jobSlotIDs))],
			userIDs[rand.Intn(len(userIDs))],
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

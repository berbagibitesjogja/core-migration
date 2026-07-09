package main

import (
	"database/sql"
	"math/rand"
	"github.com/brianvoe/gofakeit/v7"
)

func SeedAttendances(tx *sql.Tx, presenceIDs []int, jobSlotApplicantIDs []int, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO attendances (presence_id, job_slot_applicant_id, distance) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	if len(presenceIDs) == 0 || len(jobSlotApplicantIDs) == 0 {
		return ids, nil
	}

	for i := 0; i < count; i++ {
		var id int
		err = stmt.QueryRow(
			presenceIDs[rand.Intn(len(presenceIDs))],
			jobSlotApplicantIDs[rand.Intn(len(jobSlotApplicantIDs))],
			gofakeit.Number(1, 100),
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

package main

import (
	"database/sql"
	"math/rand"
	"github.com/brianvoe/gofakeit/v7"
)

func SeedCancellations(tx *sql.Tx, userIDs []int, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO cancellations (user_id, banned, tries) VALUES ($1, $2, $3) RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for i := 0; i < count; i++ {
		var id int
		err = stmt.QueryRow(
			userIDs[rand.Intn(len(userIDs))],
			gofakeit.Date(),
			gofakeit.Number(1, 5),
		).Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

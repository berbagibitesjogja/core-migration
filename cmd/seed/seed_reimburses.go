package main

import (
	"database/sql"
	"math/rand"
	"github.com/brianvoe/gofakeit/v7"
)

func SeedReimburses(tx *sql.Tx, userIDs []int, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO reimburses (user_id, method, target, amount, is_done, file) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for i := 0; i < count; i++ {
		var id int
		err = stmt.QueryRow(
			userIDs[rand.Intn(len(userIDs))],
			gofakeit.Word(),
			gofakeit.Word(),
			gofakeit.Number(10000, 1000000),
			gofakeit.Bool(),
			gofakeit.URL(),
		).Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

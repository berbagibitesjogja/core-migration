package main

import (
	"database/sql"
	"math/rand"
	"github.com/brianvoe/gofakeit/v7"
)

func SeedUsers(tx *sql.Tx, divisionIDs []int, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO users (division_id, kratos_id, role) VALUES ($1, $2, $3) RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	roles := []string{"super", "core", "staff", "member"}

	for i := 0; i < count; i++ {
		var id int
		err = stmt.QueryRow(
			divisionIDs[rand.Intn(len(divisionIDs))],
			gofakeit.UUID(),
			roles[rand.Intn(len(roles))],
		).Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

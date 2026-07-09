package main

import (
	"database/sql"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
)

func SeedDivisions(tx *sql.Tx, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO divisions (name) VALUES ($1) RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for i := 0; i < count; i++ {
		var id int
		err = stmt.QueryRow(fmt.Sprintf("%s Division", gofakeit.Word())).Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

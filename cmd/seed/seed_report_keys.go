package main

import (
	"database/sql"
	"github.com/brianvoe/gofakeit/v7"
)

func SeedReportKeys(tx *sql.Tx, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO report_keys (filename, code) VALUES ($1, $2) RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for i := 0; i < count; i++ {
		var id int
		err = stmt.QueryRow(
			gofakeit.Word()+".pdf",
			gofakeit.UUID(),
		).Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

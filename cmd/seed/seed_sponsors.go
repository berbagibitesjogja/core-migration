package main

import (
	"database/sql"
	"math/rand"
	"github.com/brianvoe/gofakeit/v7"
)

func SeedSponsors(tx *sql.Tx, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO sponsors (name, variant, address, email, phone, logo, hidden) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	variants := []string{"company", "individual"}

	for i := 0; i < count; i++ {
		var id int
		err = stmt.QueryRow(
			gofakeit.Company(),
			variants[rand.Intn(len(variants))],
			gofakeit.Address().Address,
			gofakeit.Email(),
			gofakeit.Phone(),
			gofakeit.URL(),
			gofakeit.Bool(),
		).Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

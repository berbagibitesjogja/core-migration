package main

import (
	"database/sql"
	"math/rand"
	"github.com/brianvoe/gofakeit/v7"
)

func SeedDonations(tx *sql.Tx, sponsorIDs []int, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO donations (sponsor_id, quota, remain, take, location, maps, message, media, status, reported, charity) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	statuses := []string{"active", "end"}

	for i := 0; i < count; i++ {
		var id int
		quota := gofakeit.Number(10, 100)
		remain := rand.Intn(quota + 1)
		err = stmt.QueryRow(
			sponsorIDs[rand.Intn(len(sponsorIDs))],
			quota,
			remain,
			gofakeit.Date(),
			gofakeit.Address().Address,
			"https://maps.app.goo.gl/"+gofakeit.Word(),
			gofakeit.Sentence(10),
			gofakeit.URL(),
			statuses[rand.Intn(len(statuses))],
			gofakeit.Bool(),
			gofakeit.Bool(),
		).Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

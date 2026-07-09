package main

import (
	"database/sql"
	"math/rand"
	"github.com/brianvoe/gofakeit/v7"
)

func SeedNotifies(tx *sql.Tx, heroIDs []int, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO notifies (hero_id, notify_remains) VALUES ($1, $2) ON CONFLICT (hero_id) DO NOTHING RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	
	if count > len(heroIDs) {
		count = len(heroIDs)
	}
	
	shuffled := make([]int, len(heroIDs))
	copy(shuffled, heroIDs)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	for i := 0; i < count; i++ {
		var id int
		err = stmt.QueryRow(
			shuffled[i],
			gofakeit.Number(1, 10),
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

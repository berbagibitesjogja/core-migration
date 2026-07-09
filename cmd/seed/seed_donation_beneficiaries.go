package main

import (
	"database/sql"
	"math/rand"
)

func SeedDonationBeneficiaries(tx *sql.Tx, donationIDs []int, beneficiaryIDs []int, count int) ([]int, error) {
	var ids []int
	stmt, err := tx.Prepare(`INSERT INTO donation_beneficiaries (donation_id, beneficiary_id) VALUES ($1, $2) ON CONFLICT DO NOTHING RETURNING id;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for i := 0; i < count; i++ {
		var id int
		err = stmt.QueryRow(
			donationIDs[rand.Intn(len(donationIDs))],
			beneficiaryIDs[rand.Intn(len(beneficiaryIDs))],
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

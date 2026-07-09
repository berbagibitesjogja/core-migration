package main

import (
	"database/sql"
)

func SeedAppConfiguration(tx *sql.Tx) error {
	stmt, err := tx.Prepare(`INSERT INTO app_configuration (key, value) VALUES ($1, $2);`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	configs := map[string]string{
		"maintenance_mode": "false",
		"contact_email":    "contact@berbagibites.com",
		"max_file_size":    "5242880",
	}

	for k, v := range configs {
		_, err = stmt.Exec(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	if os.Getenv("APP_ENV") == "production" {
		log.Fatal("🛑 SEEDER BLOCKED: Attempted run on production environment.")
	}

	connStr := "postgres://postgres:postgres@127.0.0.1:5432/bbj_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Database connection failure: %v", err)
	}
	defer db.Close()

	log.Println("🌱 Starting dynamic Go Faker data initialization...")

	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Failed to open db transaction: %v", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			log.Fatalf("Panic during seeding: %v", p)
		}
	}()

	// 1. Level 1
	divIDs, err := SeedDivisions(tx, 5)
	if err != nil {
		log.Fatalf("SeedDivisions error: %v", err)
	}

	sponsorIDs, err := SeedSponsors(tx, 10)
	if err != nil {
		log.Fatalf("SeedSponsors error: %v", err)
	}

	affIDs, err := SeedAffiliations(tx, 5)
	if err != nil {
		log.Fatalf("SeedAffiliations error: %v", err)
	}

	_, err = SeedReportKeys(tx, 10)
	if err != nil {
		log.Fatalf("SeedReportKeys error: %v", err)
	}

	err = SeedAppConfiguration(tx)
	if err != nil {
		log.Fatalf("SeedAppConfiguration error: %v", err)
	}

	log.Println("✅ Level 1 seeders completed.")

	// 2. Level 2
	userIDs, err := SeedUsers(tx, divIDs, 30)
	if err != nil {
		log.Fatalf("SeedUsers error: %v", err)
	}

	beneficiaryIDs, err := SeedBeneficiaries(tx, affIDs, 15)
	if err != nil {
		log.Fatalf("SeedBeneficiaries error: %v", err)
	}

	jobScheduleIDs, err := SeedJobSchedules(tx, sponsorIDs, 15)
	if err != nil {
		log.Fatalf("SeedJobSchedules error: %v", err)
	}

	donationIDs, err := SeedDonations(tx, sponsorIDs, 20)
	if err != nil {
		log.Fatalf("SeedDonations error: %v", err)
	}

	log.Println("✅ Level 2 seeders completed.")

	// 3. Level 3
	heroIDs, err := SeedHeroes(tx, beneficiaryIDs, 20)
	if err != nil {
		log.Fatalf("SeedHeroes error: %v", err)
	}

	jobSlotIDs, err := SeedJobSlots(tx, jobScheduleIDs, 25)
	if err != nil {
		log.Fatalf("SeedJobSlots error: %v", err)
	}

	_, err = SeedReimburses(tx, userIDs, 15)
	if err != nil {
		log.Fatalf("SeedReimburses error: %v", err)
	}

	_, err = SeedCancellations(tx, userIDs, 10)
	if err != nil {
		log.Fatalf("SeedCancellations error: %v", err)
	}

	_, err = SeedFoods(tx, donationIDs, 30)
	if err != nil {
		log.Fatalf("SeedFoods error: %v", err)
	}

	log.Println("✅ Level 3 seeders completed.")

	// 4. Level 4
	_, err = SeedNotifies(tx, heroIDs, 15)
	if err != nil {
		log.Fatalf("SeedNotifies error: %v", err)
	}

	_, err = SeedDonationBeneficiaries(tx, donationIDs, beneficiaryIDs, 30)
	if err != nil {
		log.Fatalf("SeedDonationBeneficiaries error: %v", err)
	}

	_, err = SeedHeroActions(tx, heroIDs, donationIDs, 30)
	if err != nil {
		log.Fatalf("SeedHeroActions error: %v", err)
	}

	jobSlotApplicantIDs, err := SeedJobSlotApplicants(tx, jobSlotIDs, userIDs, 40)
	if err != nil {
		log.Fatalf("SeedJobSlotApplicants error: %v", err)
	}

	_, err = SeedJobSlotDivisions(tx, jobSlotIDs, divIDs, 20)
	if err != nil {
		log.Fatalf("SeedJobSlotDivisions error: %v", err)
	}

	presenceIDs, err := SeedPresences(tx, jobSlotIDs, 30)
	if err != nil {
		log.Fatalf("SeedPresences error: %v", err)
	}

	log.Println("✅ Level 4 seeders completed.")

	// 5. Level 5
	_, err = SeedAttendances(tx, presenceIDs, jobSlotApplicantIDs, 40)
	if err != nil {
		log.Fatalf("SeedAttendances error: %v", err)
	}

	log.Println("✅ Level 5 seeders completed.")

	if err := tx.Commit(); err != nil {
		log.Fatalf("Transaction commit failed: %v", err)
	}

	log.Println("✅ All testing sandbox seeded successfully!")
}

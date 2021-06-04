package main

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"webapi/pkg/infra/env"
)

func main() {
	env.ConfigEnvs()
	log.Println("****************************************")
	log.Println("Migrating...")

	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("DB_HOST"), viper.GetInt("DB_PORT"), viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_NAME"),
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := createUsersTable(db); err != nil {
		log.Fatal("Error when try to create Users Table")
	}
	log.Println("****************************************")
	log.Println("Migration run successfully")
}

func createUsersTable(db *sql.DB) error {
	log.Println("****************************************")
	log.Println("Creating Users Table")

	_, err := db.Exec(`
		CREATE TABLE users
		(
			id SERIAL NOT NULL PRIMARY KEY,
			first_name VARCHAR(255) NOT NULL,
			last_name  VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP WITH TIME ZONE
		)
	`)

	if err.Error() == "pq: relation \"users\" already exists" {
		log.Println("User Table already exist")
		return nil
	}

	return err
}

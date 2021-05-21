package infra

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func GetConnection(host string, port int, user, password, dbName string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println("DB - GetConnection - ", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("DB - GetConnection - ", err)
		return nil, err
	}

	return db, nil
}

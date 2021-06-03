package infra

import (
	"database/sql"
	"fmt"
	"log"

	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/pq"
)

func GetConnection(host string, port int, user, password, dbName string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	db, err := apmsql.Open("postgres", connectionString)
	if err != nil {
		log.Printf("error while connect to database: %v", err)
		return nil, err
	}

	// err = db.Ping()
	// if err != nil {
	// 	log.Printf("error while check database connection: %v", err)
	// 	return nil, err
	// }

	return db, nil
}

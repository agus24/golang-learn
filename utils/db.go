package utils

import (
	"database/sql"
	"log"
)

func StartTransaction(db *sql.DB) (*sql.Tx, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	defer rollbackOnError(tx)

	return tx, nil
}

func rollbackOnError(tx *sql.Tx) {
	if v := recover(); v != nil {
		tx.Rollback()
		log.Println("Rollback called")
		panic(v)
	}
}

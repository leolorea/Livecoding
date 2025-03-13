package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createDB() error {

	db, err := sql.Open("sqlite3", "./image.db")
	if err != nil {
		return err
	}

	defer db.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS metadata (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		data BLOB
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	return nil

}

func InsertFile(meta MetaData) {

	db, err := sql.Open("sqlite3", "./image.db")
	if err != nil {
		log.Fatal(err)
	}

	// Insert the image data into the database
	insertSQL := `INSERT INTO metadata (name, data) VALUES (?, ?)`
	_, err = db.Exec(insertSQL, meta.Name, meta.Data)
	if err != nil {
		log.Fatal(err)
	}
}

func getFiles() []MetaData {

	db, err := sql.Open("sqlite3", "./image.db")
	if err != nil {
		log.Fatal(err)
	}

	// Insert the image data into the database
	selectSQL := `SELECT id, name, data FROM metadata`
	rows, err := db.Query(selectSQL)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var metas []MetaData

	// Iterate over the result set
	for rows.Next() {
		var meta MetaData

		// Scan the result into variables
		err := rows.Scan(&meta.Id, &meta.Name, &meta.Data)
		if err != nil {
			log.Fatal(err)
		}
		metas = append(metas, meta)

	}
	return metas
}

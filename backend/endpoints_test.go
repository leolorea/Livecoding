package main

import (
	"database/sql"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCreateDB(t *testing.T) {

	db, err := sql.Open("sqlite3", "image.db")
	if err != nil {
		t.Fatal("Error opening database:", err)
	}
	defer db.Close()

	err = createDB()
	if err != nil {
		t.Fatalf("Error creating database: %v", err)
	}

	var result int
	err = db.QueryRow("SELECT 1").Scan(&result)
	if err != nil {
		t.Fatalf("Error while consulting database: %v", err)
	}

	assert.Equal(t, 1, result)
}

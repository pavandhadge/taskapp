package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/tursodatabase/go-libsql"
)

func NewDb(dbUrl string, dbToken string, maxopenConn int, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	dbStr := dbUrl + "?authToken=" + dbToken
	db, err := sql.Open("libsql", dbStr)
	if err != nil {
		return nil, err
	}

	if maxopenConn > 0 {
		db.SetMaxOpenConns(maxopenConn) // ✅ Corrected
	}
	db.SetMaxIdleConns(maxIdleConns)

	duration, err := time.ParseDuration(maxIdleTime)
	if err == nil {
		db.SetConnMaxIdleTime(duration)
	} else {
		fmt.Println("Warning: Invalid maxIdleTime value, skipping ConnMaxIdleTime setting.")
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("Error while connecting to the database: %v\n", err) // ✅ Improved error message
		return nil, err
	}

	if err := createTable(db); err != nil { // ✅ Handle table creation errors
		return nil, fmt.Errorf("failed to initialize database schema: %w", err)
	}

	return db, nil
}

func createTable(db *sql.DB) error {
	createTaskTable := `
    CREATE TABLE IF NOT EXISTS task_lists (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        completed INTEGER DEFAULT 0,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

	_, err := db.Exec(createTaskTable)
	if err != nil {
		return fmt.Errorf("error while creating table: %v", err) // ✅ Return error
	}
	return nil
}

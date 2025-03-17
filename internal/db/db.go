package db

import (
	"database/sql"
	"fmt"
	"log"
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
		db.SetMaxIdleConns(maxopenConn)
	}
	db.SetMaxIdleConns(maxIdleConns)
	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		db.SetConnMaxIdleTime(duration)
	}
	if err := db.Ping(); err != nil {
		fmt.Printf("got and error whaile connecting the db . the erro is %v", err)
		return nil, err
	}
	createTable(db)
	return db, nil
}

func createTable(db *sql.DB) {
	createTaskTable := `
	CREATE TABLE IF NOT EXISTS task_list (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    completed INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
	`
	_, err := db.Exec(createTaskTable)
	if err != nil {
		log.Fatal("error while making table ", err)
	}
}

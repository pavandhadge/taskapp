package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pavandhadge/taskapp/internal/db"
	"github.com/pavandhadge/taskapp/internal/repository"
)

func main() {
	err := godotenv.Load()
	dbConfig := dbconfig{
		dbUrl:        os.Getenv("DB_URL"),
		dbToken:      os.Getenv("DB_TOKEN"),
		maxopenConn:  5,
		maxIdleConns: 2,
		maxIdleTime:  "15m",
	}
	fmt.Print(dbConfig)
	db, err := db.NewDb(dbConfig.dbUrl, dbConfig.dbToken, dbConfig.maxopenConn, dbConfig.maxIdleConns, dbConfig.maxIdleTime)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	print("database connection pool established")
	app := &application{
		config: dbConfig,
		addr:   os.Getenv("PORT"),
		store:  repository.NewStorage(db),
	}
	mux := app.mount()
	if err := app.run(mux); err != nil {
		fmt.Print("err while starting server ", err)
	}
}

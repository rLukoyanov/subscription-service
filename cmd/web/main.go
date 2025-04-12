package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

func main() {
	// init db
	db := initDB()
	db.Ping()
	// create sessions
	// init chanels
	// init wg
	// init cfg
	// set up sending emails
	// listen for web conns
}

func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("cant connect to db")
	}

	return conn
}

func connectToDB() *sql.DB {
	counts := 0
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("postgres not yet ready")
		} else {
			log.Println("connected to db")
			return connection
		}

		if counts > 10 {
			return nil
		}

		counts++
		time.Sleep(1 * time.Second)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

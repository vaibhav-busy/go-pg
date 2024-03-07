package db

import (
	"log"
	"os"
	"time"

	pg "github.com/go-pg/pg"
)

func Connect() *pg.DB {

	opts := &pg.Options{
		User:         "postgres",
		Password:     "5679",
		Addr:         "localhost:5432",
		Database:     "mydb",
		DialTimeout:  30 * time.Second,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		IdleTimeout:  30 * time.Minute,
		MaxConnAge:   1 * time.Minute,
		PoolSize:     20,
	}

	db := pg.Connect(opts)

	if db == nil {
		log.Printf("Error connecting to database")
		os.Exit(100)
	}

	log.Printf("Successfully connected to database")

	CreateProdItemsTable(db)

	// defer db.Close()

	// closeErr := db.Close()

	// if closeErr != nil {
	// 	log.Printf("Error while closing database %v", closeErr)
	// 	os.Exit(100)
	// }

	log.Printf("Connection closed successfully")
	return db
}

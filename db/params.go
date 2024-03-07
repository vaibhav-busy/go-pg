package db

import (
	"log"

	"github.com/go-pg/pg"
)

type Student struct{
	Name string
	Addr string
}


func PlaceHolder(db *pg.DB) error {
	// var value int
	var value string

	stu:=Student{
		Name: "Vaibhav",
		Addr:"10-A kp" ,
	}

	// var query string = "SELECT ?0"
	// _, selectErr := db.Query(pg.Scan(&value), query, 3,4,5)

	var query string = "SELECT ?addr"
	_, selectErr := db.Query(pg.Scan(&value), query, stu)


	if selectErr != nil {
		log.Printf("Error while running the select query")
		return selectErr
	}

	log.Printf("Scan successful, scanned values:  %v",value)
	return nil
}

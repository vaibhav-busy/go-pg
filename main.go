package main

import (
	db "go-pg/db"
	"log"
	"time"

	pg "github.com/go-pg/pg"
)

func main() {
	log.Printf("Hello ")
	pg_db := db.Connect()
	// SaveProduct(pg_db)
	// db.PlaceHolder(pg_db)
	// DeleteItem(pg_db)
	UpdateItem(pg_db)
	// GetById(pg_db)
	pg_db.Close()
}

func SaveProduct(dbRef *pg.DB) {

	newProd1 := &db.ProductItem{
		Name:  "Prod 9",
		Price: 13.123456789,
		Feature: struct {
			Name string
			Desc string
			Imp  int
		}{
			Name: "F1",
			Desc: "F1 desc",
			Imp:  234,
		},

		CreatedAt: time.Now(),
		IsActive:  true,
	}

	newProd2 := &db.ProductItem{
		Name:  "Prod 10",
		Price: 13.123456789,
		Feature: struct {
			Name string
			Desc string
			Imp  int
		}{
			Name: "F1",
			Desc: "F1 desc",
			Imp:  234,
		},

		CreatedAt: time.Now(),
		IsActive:  true,
	}

	totalItems := []*db.ProductItem{newProd1, newProd2}
	newProd1.SaveMultiple(dbRef, totalItems)
}

func DeleteItem(dbRef *pg.DB) {

	newPI := &db.ProductItem{
		Name: "Prod 8",
	}

	newPI.DeleteItem(dbRef)

}

func UpdateItem(dbRef *pg.DB) {

	newPI := &db.ProductItem{
		ID:    3,
		Price: 400,
		IsActive: false,
	}

	newPI.UpdateItem(dbRef)
}

func GetById(dbRef *pg.DB) {

	// newPI := &db.ProductItem{
	// 	ID: 1,
	// }

	newPI := &db.ProductItem{
		ID: 1,
		Name: "Prod 1",
	}
	newPI.GetById(dbRef)
}
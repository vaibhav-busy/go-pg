package db

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type ProductItem struct {
	tablename  struct{} `sql:"product_items_collection"`
	Refpointer int      `sql:"-"`

	ID      int     `sql:"id,pk`
	Name    string  `sql:"name,unique"`
	Price   float64 `sql:"price, type:real"`
	Feature struct {
		Name string
		Desc string
		Imp  int
	}

	CreatedAt time.Time `sql:"created_at"`
	IsActive  bool      `sql:"is_active"`
}

func CreateProdItemsTable(db *pg.DB) error {

	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := db.CreateTable(&ProductItem{}, opts)

	if createErr != nil {
		log.Printf("Error in table creation %v", createErr)
	}

	return nil

}

func (pi *ProductItem) Save(db *pg.DB) error {

	insertErr := db.Insert(pi)

	if insertErr != nil {
		log.Printf("Error while inserting value to db %v", insertErr)
		return insertErr
	}

	log.Printf("Item successfully inserted", pi.Name)

	return nil

}

func (pi *ProductItem) SaveAndReturn(db *pg.DB) (*ProductItem, error) {

	InsertResult, insertErr := db.Model(pi).Returning("*").Insert()
	if insertErr != nil {
		log.Printf("Error while inserting new item %v", insertErr)
	}

	log.Printf("Item Inserted Successfully")
	log.Printf("Received new value, result is : %v", InsertResult.RowsAffected())
	return pi, nil
}

func (pi *ProductItem) SaveMultiple(db *pg.DB, items []*ProductItem) error {

	_, insertErr := db.Model(items[0], items[1]).Insert()

	if insertErr != nil {
		log.Printf("Error while inserting multiple values %v", insertErr)
		return insertErr
	}

	log.Printf("Bulk insert Successful")
	return nil

}

func (pi *ProductItem) DeleteItem(db *pg.DB) error {

	_, deleteErr := db.Model(pi).Where("name=?name").Delete()

	if deleteErr != nil {
		log.Printf("Error in deleting value %v", deleteErr)
		return deleteErr
	}

	log.Printf("Successfully deleted item :  %v", pi.Name)
	return nil

}

func (pi *ProductItem) UpdateItem(db *pg.DB) error {

	tx, txErr := db.Begin()
	if txErr != nil {
		log.Printf("Error while opening transaction %v", txErr)
		return txErr
	}

	_, updateErr := tx.Model(pi).Set("price=?price").Where("id=?id").Update()

	if updateErr != nil {
		log.Printf("Error in updating value %v", updateErr)
		tx.Rollback()
		return updateErr
	}


	_, updateErr2 := tx.Model(pi).Set("is_acti=?0",false).Where("id=?id").Update()

	if updateErr2 != nil {
		log.Printf("Error in updating value %v", updateErr2)
		tx.Rollback()
		return updateErr2
	}

	tx.Commit()

	log.Printf("Successfully updated item :  %v", pi.Name)
	return nil

}

func (pi *ProductItem) GetById(db *pg.DB) error {

	// getErr:=db.Select(pi)
	// getErr := db.Model(pi).Where("id=?0",pi.ID).Select()
	// getErr := db.Model(pi).Column("name","price").Where("id=?0", pi.ID).Select()

	// getErr := db.Model(pi).Column("name","price").
	// Where("id=?0", pi.ID).
	// Where("name=?0",pi.Name).
	// Select()

	getErr := db.Model(pi).Column("name", "price").
		Where("id=?0", pi.ID).
		WhereOr("id=?0", 3).
		Offset(1).
		Limit(1).
		Order("id desc").
		Select()

	if getErr != nil {
		log.Printf("Error in getting value %v", getErr)
		return getErr
	}

	log.Printf("Get by id successful, item is : %v", *pi)
	return nil

}

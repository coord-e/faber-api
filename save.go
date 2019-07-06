package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Entry struct {
	ID string `gorm:"UNIQUE"`
	Options Options
	Result Result
}

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Entry{})
	return db, nil
}

func Save(db *gorm.DB, opts Options, res *Result) (string, error) {
	id := uuid.New().String()
	entry := Entry{
		ID: id,
		Options: opts,
		Result: *res,
	}
	db.Create(&entry)
	return id, nil
}

func Restore(db *gorm.DB, id string) (Options, Result, error) {
	entry := Entry{}
	if db.Where("id = ?", id).First(&entry).RecordNotFound() {
		err := fmt.Errorf("can't find a saved entry with id %s", id)
		return Options{}, Result{}, err
	}
	return entry.Options, entry.Result, nil
}

package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func New(ds string) (*gorm.DB, func()) {
	db, err := gorm.Open("postgres", ds)
	if err != nil {
		log.Println("Fail to open db")
		panic(err)
	}

	return db, func() {
		err := db.Close()
		if err != nil {
			log.Println("Failed to close DB by error", err)
		}
	}
}

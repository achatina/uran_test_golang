package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func Get() *gorm.DB {
	return db
}

func Open() {
	var err error
	db, err = gorm.Open("mysql", "root:password@/uran")

	if err != nil {
		log.Fatal(err)
	}
}

func Close() {
	db.Close()
}

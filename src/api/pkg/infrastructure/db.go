package infrastructure

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(dsn string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	count := 5
	for count > 1 {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {

			time.Sleep(2 * time.Second)
			count--
			log.Printf("retry... count:%v\n", count)
			continue
		}
		break
	}

	return db, err
}

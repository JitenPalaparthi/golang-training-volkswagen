package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MAXRETRIES    = 5
	RETRYDURATION = time.Second * 5
)

func GetDatabase(dsn string) (*gorm.DB, error) {
	counter := 1
retry:
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		counter++
		time.Sleep(RETRYDURATION)
		if counter > MAXRETRIES {
			return nil, fmt.Errorf("retried for %d number of times.Error: %s", counter, err.Error())
		}
		log.Println("retrying to connect to the database.Current retry is", counter, "times")
		goto retry
	}
	return db, nil
}

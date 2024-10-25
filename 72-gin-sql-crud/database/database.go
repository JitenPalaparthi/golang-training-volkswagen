package database

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

// var Pool []*sql.DB
// var PoolCount uint8 = 5

const (
	USERTABLE = `
	CREATE TABLE IF NOT EXISTS users(
	id INT AUTO_INCREMENT,
	name VARCHAR(100) NOT NULL,
	email VARCHAR(100) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id));`
)

func GetDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, err
}

func CreateTable(db *sql.DB, script string) error {
	if db == nil {
		return errors.New("invalid db connection")
	}
	_, err := db.Exec(script)
	if err != nil {
		return err
	}
	return nil
}

func AutoCreateTables(db *sql.DB) error {
	if db == nil {
		return errors.New("invalid db connection")
	}
	return CreateTable(db, USERTABLE)
}

func CloseDatabase(db *sql.DB) error {
	return db.Close()
}

func GetDBConnectionFromPool(pool []*sql.DB) (*sql.DB, error) {
	// what algorthm you implement here to get the connection from the pool is imp
	return nil, nil
}

func GetDatabasePool(dsn string, poolCount uint8) (Pool []*sql.DB, err error) {
	for i := 1; i <= int(poolCount); i++ {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, err
		}
		Pool = append(Pool, db)
	}
	return Pool, nil
}

func CloseDatabasePool(pool []*sql.DB) error {
	for _, p := range pool {
		err := p.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

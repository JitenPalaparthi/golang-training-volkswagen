package database

import (
	"database/sql"
	"demo/models"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type UserDB struct {
	DB *sql.DB
}

func (ud *UserDB) Insert(user models.User) (int64, error) {
	if user.Name == "" || user.Email == "" {
		return 0, errors.New("invalid name or email")
	}
	if ud.DB == nil {
		return 0, errors.New("invalid db connection")
	}

	insertUser := `INSERT INTO users(name,email)VALUES(?,?)`
	r, err := ud.DB.Exec(insertUser, user.Name, user.Email)
	if err != nil {
		return 0, err
	}
	return r.LastInsertId()
}

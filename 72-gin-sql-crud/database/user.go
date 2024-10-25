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

func (ud *UserDB) GetUser(id int) (*models.User, error) {
	user := new(models.User)
	seletUser := `SELECT id,name,email,created_at FROM users where id=?`
	row := ud.DB.QueryRow(seletUser, id)
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt); err != nil {
		return user, err
	}
	return user, nil
}

func (ud *UserDB) GetAllUsers() ([]models.User, error) {
	users := make([]models.User, 0)
	seletUser := `SELECT id,name,email,created_at FROM users`
	rows, err := ud.DB.Query(seletUser)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (ud *UserDB) GetAllUsersBy(limit, offset int) ([]models.User, error) {
	users := make([]models.User, 0)
	seletUsers := `SELECT id,name,email,created_at FROM users ORDER BY id LIMIT ? OFFSET ?`
	rows, err := ud.DB.Query(seletUsers, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (ud *UserDB) DeleteUser(id int) (int64, error) {
	seletUser := `DELETE FROM users where id=?`
	r, err := ud.DB.Exec(seletUser, id)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

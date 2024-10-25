package database

import (
	"demo/models"
	"time"

	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func (ud *UserDB) Insert(user models.User) (int64, error) {
	err := ud.DB.AutoMigrate(models.User{})
	if err != nil {
		return 0, err
	}
	user.CreatedAt = uint64(time.Now().Unix())
	tx := ud.DB.Create(&user)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int64(user.Id), nil
}

func (ud *UserDB) GetUser(id int) (*models.User, error) {
	user := new(models.User)
	tx := ud.DB.First(user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	//row := ud.DB.Table("users").Select("id,name,email,created_at").Joins().Where("id=?", id).Row()

	return user, nil
}

// func (ud *UserDB) GetAllUsers() ([]models.User, error) {
// 	users := make([]models.User, 0)
// 	seletUser := `SELECT id,name,email,created_at FROM users`
// 	rows, err := ud.DB.Query(seletUser)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for rows.Next() {
// 		user := models.User{}
// 		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt); err != nil {
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}
// 	return users, nil
// }

// func (ud *UserDB) GetAllUsersBy(limit, offset int) ([]models.User, error) {
// 	users := make([]models.User, 0)
// 	seletUsers := `SELECT id,name,email,created_at FROM users ORDER BY id LIMIT ? OFFSET ?`
// 	rows, err := ud.DB.Query(seletUsers, limit, offset)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for rows.Next() {
// 		user := models.User{}
// 		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt); err != nil {
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}
// 	return users, nil
// }

func (ud *UserDB) DeleteUser(id int) (int64, error) {
	tx := ud.DB.Delete(&models.User{}, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

// func (ud *UserDB) UpdateUserQ(id int, data map[string]string) (int64, error) {
// 	updateQuery := "UPDATE users SET "
// 	query := ""
// 	for k, v := range data {
// 		query += k + "=" + v + ","
// 	}
// 	updateQuery += query[:len(query)-1] + " WHERE id=?;"
// 	r, err := ud.DB.Exec(updateQuery, id)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return r.RowsAffected()
// }

// func (ud *UserDB) UpdateUser(id int, user models.User) (int64, error) {
// 	actualUser, err := ud.GetUser(id)

// 	if err != nil {
// 		return 0, err
// 	}
// 	if user.Email == "" {
// 		user.Email = actualUser.Email
// 	}
// 	if user.Name == "" {
// 		user.Name = actualUser.Name
// 	}

// 	updateQuery := "UPDATE users SET name=?,email=? WHERE id=?;"
// 	r, err := ud.DB.Exec(updateQuery, user.Name, user.Email, id)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return r.RowsAffected()
// 	return 1, nil
// }

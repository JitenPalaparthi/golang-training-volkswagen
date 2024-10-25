package models

import (
	"encoding/json"
)

type User struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
	//Address   []UserAddress `gorm:foriegnkey:userID`
	CreatedAt uint64 `json:"created_at" gorm:"column:created_at"`
}

type UserAddress struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Address   string `json:"address"`
	UserID    int    `json:"userId"`
	CreatedAt uint64 `json:"created_at" gorm:"column:created_at"`
}

func (u *User) ToString() string {
	bytes, err := json.Marshal(u) // handle this error
	if err != nil {
		return ""
	}
	return string(bytes)
}

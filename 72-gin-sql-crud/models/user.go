package models

import "encoding/json"

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func (u *User) ToString() string {
	bytes, err := json.Marshal(u) // handle this error
	if err != nil {
		return ""
	}
	return string(bytes)
}

package interfaces

import "demo/models"

type IUser interface {
	Insert(user models.User) (int64, error)
	GetUser(id int) (*models.User, error)
	// GetAllUsers() ([]models.User, error)
	// GetAllUsersBy(limit, offset int) ([]models.User, error)
	DeleteUser(id int) (int64, error)
	// UpdateUser(id int, user models.User) (int64, error)
}

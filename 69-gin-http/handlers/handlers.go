package handlers

import "demo/models"

var (
	ChStudent chan models.Student
)

func init() {
	ChStudent = make(chan models.Student, 5)
}

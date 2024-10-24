package handlers

import (
	"demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Studenthandler struct {
	//ChStudent chan models.Student
}

func (s *Studenthandler) Create(ctx *gin.Context) {
	student := new(models.Student)
	err := ctx.Bind(student)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}
	// if s.ChStudent != nil {
	// 	s.ChStudent <- *student
	// }
	ChStudent <- *student
	ctx.JSON(http.StatusOK, gin.H{
		"message": "student successfully created",
	})

}

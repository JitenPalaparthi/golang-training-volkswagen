package handlers

import (
	"demo/interfaces"
	"demo/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	IUser interfaces.IUser
}

func (uh *UserHandler) Create(ctx *gin.Context) {
	user := new(models.User)

	err := ctx.Bind(user)
	//
	// bytes, _ := io.ReadAll(ctx.Request.Body)
	// json.Unmarshal(bytes, user)

	// decoder := json.NewDecoder(ctx.Request.Body)
	// decoder.Decode(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}
	n, err := uh.IUser.Insert(*user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not insert user"})
		ctx.Abort()
		return
	}
	ChAudit <- models.Audit{User: "system", LastModified: time.Now().Unix(), Data: user.ToString(), Action: "insert"}
	ctx.String(201, fmt.Sprint("user id:", n))
}

func (us *UserHandler) GetUser(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		ctx.Abort()
		return
	}
	id1, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		ctx.Abort()
		return
	}

	user, err := us.IUser.GetUser(id1)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	ChAudit <- models.Audit{User: "system", LastModified: time.Now().Unix(), Data: id, Action: "get"}

	ctx.JSON(200, user)

}

// func (us *UserHandler) GetAllUsers(ctx *gin.Context) {
// 	users, err := us.IUser.GetAllUsers()
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		ctx.Abort()
// 		return
// 	}
// 	ctx.JSON(200, users)
// }

// func (us *UserHandler) GetAllUsersBy(ctx *gin.Context) {

// 	limit, ok := ctx.Params.Get("l")
// 	if !ok {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid limit"})
// 		ctx.Abort()
// 		return
// 	}
// 	l1, err := strconv.Atoi(limit)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid limit"})
// 		ctx.Abort()
// 		return
// 	}

// 	offset, ok := ctx.Params.Get("f")
// 	if !ok {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid offset"})
// 		ctx.Abort()
// 		return
// 	}
// 	o1, err := strconv.Atoi(offset)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid offset"})
// 		ctx.Abort()
// 		return
// 	}

// 	users, err := us.IUser.GetAllUsersBy(l1, o1)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		ctx.Abort()
// 		return
// 	}
// 	ctx.JSON(200, users)
// }

func (us *UserHandler) DeleteUser(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		ctx.Abort()
		return
	}
	id1, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		ctx.Abort()
		return
	}

	r, err := us.IUser.DeleteUser(id1)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	ChAudit <- models.Audit{User: "system", LastModified: time.Now().Unix(), Data: id, Action: "delete"}

	ctx.String(200, fmt.Sprintf("Records Deleted:%d", r))

}

// func (us *UserHandler) UpdateUser(ctx *gin.Context) {
// 	id, ok := ctx.Params.Get("id")
// 	if !ok {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
// 		ctx.Abort()
// 		return
// 	}
// 	id1, err := strconv.Atoi(id)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
// 		ctx.Abort()
// 		return
// 	}
// 	user := new(models.User)

// 	err = ctx.Bind(user)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		ctx.Abort()
// 		return
// 	}

// 	r, err := us.IUser.UpdateUser(id1, *user)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		ctx.Abort()
// 		return
// 	}
// 	ctx.String(200, fmt.Sprintf("Records Updated:%d", r))

// }

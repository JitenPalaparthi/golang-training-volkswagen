package main

import (
	"demo/database"
	"demo/handlers"
	"demo/helpers"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
	DSN  = `admin:admin@tcp(127.0.0.1:3306)/demodb`
)

func main() {
	r := gin.Default() // default router

	db, err := database.GetDatabase(DSN)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.CloseDatabase(db)

	err = database.AutoCreateTables(db)
	if err != nil {
		log.Println(err.Error())
	}
	userDB := &database.UserDB{DB: db}
	userHandler := handlers.UserHandler{IUser: userDB}

	r.GET("/ping", handlers.Ping)
	r.GET("/health", handlers.Health)
	r.POST("/user", userHandler.Create)
	r.GET("/user/:id", userHandler.GetUser)
	r.GET("/user/all", userHandler.GetAllUsers)
	r.GET("/user/all/:l/:f", userHandler.GetAllUsersBy)
	r.PUT("/user/:id", userHandler.UpdateUser)
	r.DELETE("/user/:id", userHandler.DeleteUser)
	go helpers.DoAudit("audit/audit.log", handlers.ChAudit)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

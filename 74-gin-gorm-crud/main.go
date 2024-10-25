package main

import (
	"demo/database"
	"demo/handlers"
	"demo/helpers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
	DSN  = `admin:admin@tcp(127.0.0.1:3306)/demodb`
)

func main() {
	r := gin.Default() // default router

	r.Use(gin.Logger()) //
	r.Use(gin.Recovery())
	r.Use(CheckHeder)

	db, err := database.GetDatabase(DSN)
	if err != nil {
		log.Fatal(err.Error())
	}
	if err != nil {
		log.Println(err.Error())
	}
	userDB := &database.UserDB{DB: db}
	userHandler := handlers.UserHandler{IUser: userDB}

	r.GET("/ping", special, handlers.Ping)
	r.GET("/health", handlers.Health)
	r.POST("/user", userHandler.Create)
	r.GET("/user/:id", userHandler.GetUser)
	// r.GET("/user/all", userHandler.GetAllUsers)
	// r.GET("/user/all/:l/:f", userHandler.GetAllUsersBy)
	// r.PUT("/user/:id", userHandler.UpdateUser)
	r.DELETE("/user/:id", userHandler.DeleteUser)
	go helpers.DoAudit("audit/audit.log", handlers.ChAudit)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func CheckHeder(c *gin.Context) {
	if c.GetHeader("username") == "admin" && c.GetHeader("password") == "admin123" {
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "access denied due to bad username password"})
		c.Abort()
		return
	}
}

func special(c *gin.Context) {
	if c.GetHeader("ping") == "yes" {
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "ping header not found"})
		c.Abort()
		return
	}
}

package main

import (
	"demo/handlers"
	"demo/helpers"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// https://github.com/gin-gonic/gin
// https://gin-gonic.com/docs/examples/upload-file/

var (
	PORT string
)

func main() {
	r := gin.Default() // default router

	r.GET("/ping", handlers.Ping)
	r.GET("/health", handlers.Health)
	r.POST("/student", new(handlers.Studenthandler).Create)

	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "files/"+file.Filename)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	go helpers.CreateFileChan(handlers.ChStudent)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

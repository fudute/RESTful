package main

import (
	"net/http"

	"github.com/fudute/RESTful/httpd/handler"
	"github.com/fudute/RESTful/platform/users"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repo := users.Repo{
		Users: []users.User{},
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"resp": "pong",
		})
	})

	r.GET("/users", handler.GetUser(&repo))
	r.POST("/users", handler.AddUser(&repo))

	r.Run()
}

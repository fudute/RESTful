package handler

import (
	"net/http"

	"github.com/fudute/RESTful/platform/users"
	"github.com/gin-gonic/gin"
)

// GetUser in repo
// 使用这种first class function的方式，可以让返回的函数能够访问到GetUser的传入参数
// 可以传入数据库链接或者连接池
func GetUser(repo *users.Repo) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, repo.GetAll())
	}
}

// AddUser to repo
func AddUser(repo *users.Repo) func(c *gin.Context) {
	return func(c *gin.Context) {
		u := users.User{}
		c.Bind(&u)
		repo.Add(u)

		c.Status(http.StatusNoContent)
	}
}

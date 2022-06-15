package controller

import (
	"easy-apis/admin/logics"
	"github.com/gin-gonic/gin"
	"log"
)

func SignUp(c *gin.Context) {
	username := c.PostForm("username")
	password := logics.EncryptPassword([]byte(c.PostForm("password")))

	log.Println(username, password)
}

/*func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
}*/

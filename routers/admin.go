package routers

import (
	"easy-apis/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitAdminRouters() {
	r := gin.Default()

	// 登录
	r.POST("login", controller.Login)

}

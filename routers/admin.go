package routers

import (
	"easy-apis/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitAdminRouters() {
	r := gin.Default()

	// 用户路由组
	userGroup := r.Group("/user/")
	userGroup.Use()
	{
		userGroup.POST("signup", controller.SignUp)
		//userGroup.POST("login", controller.Login)
	}

	r.Run(":8888")
}

package routes

import (
	v1 "hello_gin/app/v1"

	"github.com/gin-gonic/gin"
)

func InitArticleRouter(Router *gin.RouterGroup) {
	groupRouter := Router.Group("article")
	//.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		groupRouter.POST("add", v1.Add) // 新增文章
		groupRouter.GET("add", v1.Add)  // 新增文章
		// groupRouter.DELETE("delete", v1.DeleteUser) //删除文章
	}
}

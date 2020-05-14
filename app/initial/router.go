package initial

import (
	"fmt"
	"hello_gin/app/global"

	"hello_gin/app/middleware"
	"hello_gin/app/routes"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	// 跨域
	Router.Use(middleware.Cors())
	fmt.Println(global.GVA_LOG)
	global.GVA_LOG.Debug("use middleware cors")
	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")
	routes.InitArticleRouter(ApiGroup) // 注册用户路由

	return Router
}

package router

import (
	"github.com/eyotang/game-proxy/server/api/v1"
	"github.com/eyotang/game-proxy/server/middleware"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base").Use(middleware.NeedInit())
	{
		BaseRouter.POST("login", v1.Login)
		BaseRouter.POST("captcha", v1.Captcha)
	}
	return BaseRouter
}

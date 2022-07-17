package router

import (
	"github.com/gin-gonic/gin"
	"go_web/interaction/controller/login"
	"go_web/interaction/controller/menu"
	"net/http"
)

func Init() *gin.Engine {
	r := gin.Default()

	// 跨域请求
	r.Use(func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		if origin != "" {
			context.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			context.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			context.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	})

	r.POST("/login", login.Cont.Login)
	r.GET("/menus", menu.Cont.GetMenus)

	userGroup := r.Group("/user")
	userRouter(userGroup)
	/** 其他 路由 */

	return r
}

package router

import (
	"github.com/gin-gonic/gin"
	"go_web/interaction/controller/user"
)

func userRouter(r *gin.RouterGroup) {
	r.GET("/page", user.Cont.GetUsers)
	r.POST("/changeState", user.Cont.ChangeState)
	r.POST("/add", user.Cont.AddUser)
	r.GET("/get", user.Cont.GetById)
	r.POST("/edit", user.Cont.EditUser)
	r.GET("/delete", user.Cont.DeleteUser)
}

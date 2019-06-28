package router

import (
	"github.com/gin-gonic/gin"
	"ukid/controller"
	"xormt"
)

func init(){
	addRegisterHandler(RegisterUserRouter)
}

func RegisterUserRouter(engine *gin.Engine){
	ctrl := &controller.UserController{}

	v1 := engine.Group("/api/v1")
	v1.GET("/users/:id", xormt.HandlerGin(ctrl.Get))
	v1.GET("/users", xormt.HandlerGin(ctrl.GetUsers))
}

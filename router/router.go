package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.GET("/", controller.GoIndex)
	e.GET("/register", controller.GoRegister)
	e.POST("/register", controller.Register)
	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)

	// 博客
	e.GET("/post_index", controller.GetPostIndex) // 获取博客列表
	e.POST("/post", controller.AddPost)           // 添加博客
	e.GET("/post", controller.GoAddPost)          // 跳转添加博客页面
	e.GET("/postDetail/:id", controller.PostDetail)

	e.Run()
}

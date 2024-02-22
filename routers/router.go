package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kkdZHC/bubble_demo/controller"
)

func SetupRouter() (r *gin.Engine) {
	r = gin.Default()
	//加载静态文件(参数1：html中访问的路径，参数2：实际的文件夹位置)
	r.Static("/static", "static")
	//加载模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	//v1
	v1Group := r.Group("v1")
	{
		//业务
		//添加
		v1Group.POST("/todo", controller.CreateTodo)
		//查看
		v1Group.GET("/todo", controller.GetTodoList)

		//修改
		v1Group.PUT("/todo/:id", controller.UpdataTodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return
}

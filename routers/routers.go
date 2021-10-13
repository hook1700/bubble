package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetRouters() *gin.Engine  {
	r := gin.Default()
	//静态变量设置
	r.Static("/static","static")
	//模板设置
	r.LoadHTMLGlob("template/*")
	r.GET("/", controller.IndexHandle)
	//v1 设置访问前缀
	v1Group := r.Group("v1")
	{
		//代办事项
		//添加
		v1Group.POST("/todo",controller.CreateTodo)
		//查看全部
		v1Group.GET("/todo", controller.GetTodoList)
		//查看某一
		//v1Group.GET("/todo/:id", controller.GetOneTodo)
		//修改
		v1Group.PUT("/todo/:id",controller.UpdateOneTodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteOneTodo)
	}
	return r
}

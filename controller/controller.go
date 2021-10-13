package controller

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandle(c *gin.Context) {
	c.HTML(http.StatusOK,"index.html",nil)
}




func CreateTodo(c *gin.Context) {
	//获取json数据
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.CreateOneTodo(&todo)
	if err !=nil {
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,todo)
	}
}

func GetTodoList(c *gin.Context) {
	 todolist, err := models.GetAllTodo()
	 if err !=nil{
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,todolist)
	}
}

//func GetOneTodo(c *gin.Context) {
//	id := c.Param("id")
//	c.JSON(http.StatusOK,gin.H{
//		"data":models.GetOneTodo(id),
//	})
//}


func UpdateOneTodo(c *gin.Context) {
	id,ok := c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusOK,gin.H{
			"err":"无效id",
		})
		return
	}
	todo,err := models.GetATodo(id)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
	})
		return
	}
	c.BindJSON(&todo)
	err = models.UpdateATodo(todo)  //?为什么么不用传指针
	if err!= nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK,todo)
	}

}

func DeleteOneTodo (c *gin.Context) {
	id,ok := c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusOK,gin.H{
			"err":"无效id",
		})
		return
	}
	err := models.DeleteATodo(id)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}
	c.JSON(http.StatusOK,gin.H{
		id:"delete",
	})
}
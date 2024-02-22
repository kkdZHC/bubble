package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kkdZHC/bubble_demo/models"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	//页面填写代办事项 点击提交 发post请求
	//1.从请求中拿出数据(参数绑定)
	var todo models.Todo
	c.ShouldBind(&todo)

	//2.存入数据库
	err := models.CreateTodo(&todo)
	//3.返回响应
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(c *gin.Context) {
	//查询todos表中的所有数据SELECT * FROM todos
	todolist, err := models.GetTodoList()
	//返回响应
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

func UpdataTodo(c *gin.Context) {
	id, _ := c.Params.Get("id")
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	err = models.UpdataTodo(id, todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	id, _ := c.Params.Get("id")
	fmt.Printf("id: %v\n", id)

	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	err = models.DeleteTodo(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

package main

import (
	"github.com/kkdZHC/bubble_demo/dao"
	"github.com/kkdZHC/bubble_demo/models"
	"github.com/kkdZHC/bubble_demo/routers"
)

func main() {
	//创建数据库 sql: create database bubble;
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	//程序退出时关闭数据库连接
	defer dao.Close()
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	//注册路由
	r := routers.SetupRouter()
	r.Run()

}

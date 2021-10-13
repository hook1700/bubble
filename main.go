package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"

)

func main() {
	//创建数据库 create database bubble

	//连接数据库
	err := dao.InitMySQL()
	if err != nil{
		panic(err)
	}
	//模型对应表
	dao.DB.AutoMigrate(&models.Todo{})
	//延迟关闭数据库
	defer dao.Close()

	//初始化
	r := routers.SetRouters()
	r.Run(":8888")







}

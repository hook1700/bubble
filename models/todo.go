package models

import "bubble/dao"

//创建TODO 结构体

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

/*
	TODO 增删改查
 */

// 增加一条数据

func CreateOneTodo(todo *Todo)  (err error){
	err = dao.DB.Create(todo).Error
	return
}

// 查询全部数据

func GetAllTodo ()  (todolist []*Todo,err error,){
	if err = dao.DB.Find(&todolist).Error;err != nil{
		return nil,err
	}
	return
}

// 查询单条数据

func GetATodo (id string)  (todo *Todo,err error,){
	todo = new(Todo)   //为什么要new?
	if err = dao.DB.Where("id=?", id).First(todo).Error;err != nil{
		return nil,err
	}
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodo (id string)  (err error,) {
	err = dao.DB.Where("id = ?", id).Delete(&Todo{}).Error
		return

}

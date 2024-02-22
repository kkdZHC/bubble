package models

import "github.com/kkdZHC/bubble_demo/dao"

// todo model
type Todo struct {
	ID     int    `json:"id`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	if err != nil {
		return err
	}
	return
}

func GetTodoList() (todolist []*Todo, err error) {
	//查询todos表中的所有数据SELECT * FROM todos
	err = dao.DB.Find(&todolist).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetATodo(id string) (todo *Todo, err error) {
	err = dao.DB.First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return
}

func UpdataTodo(id string, todo *Todo) (err error) {
	err = dao.DB.Model(Todo{}).Where("id=?", id).Update("status", !todo.Status).Error
	if err != nil {
		return err
	}
	return
}
func DeleteTodo(todo *Todo) (err error) {
	err = dao.DB.Delete(&todo).Error
	if err != nil {
		return err
	}
	return
}

package mapper

import (
	"NetworkControl/model"
	"errors"
	"gorm.io/gorm"
)

// AddTodo 创建一个记录 数据库操作
func AddTodo(title string, content string, userid uint) (uint, error) {
	var todo = model.Todo{Content: content, Title: title, UserId: userid}
	err := db.Create(&todo).Error
	if err != nil {
		return 0, err
	}
	return todo.ID, nil
}

// GetTodoByTodoId 查找记录 数据库操作
func GetTodoByTodoId(todoid uint) (*model.Todo, error) {
	var t model.Todo
	err := db.Where("user_id=?", todoid).First(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetLogin 数据库操作：验证登录
func GetLogin(login model.Login) (bool, error) {
	err := db.Where("id = ? and password = ?", login.ID, login.Password).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

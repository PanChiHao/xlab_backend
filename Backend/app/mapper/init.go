package mapper

import (
	"NetworkControl/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// Init 连接mysql
func Init(source string) error {

	// 连接上数据库
	db, err := gorm.Open(mysql.Open(source), &gorm.Config{})
	if err != nil {
		return err
	}

	// 使用 automigrate 建立todo
	err = db.AutoMigrate(&model.Todo{})
	if err != nil {
		return err
	}
	return nil
}

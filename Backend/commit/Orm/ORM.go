package main

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	Id     uint `gorm:"primaryKey"`
	Name   string
	Passwd string
}

type Todo struct {
	Id      uint `gorm:"primaryKey"`
	UserId  uint
	Title   string
	Content string
}

func main() {

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 使用 orm 连接数据库
	db, err1 := gorm.Open(mysql.Open(viper.GetString("source")), &gorm.Config{})
	if err1 != nil {
		log.Fatal(err1)
	}

	// 使用 automigrate 建立上述两张表
	err = db.AutoMigrate(&User{}, &Todo{})
	if err != nil {
		log.Fatal(err)
	}

	// 新建用户
	db.Create(&User{
		Name:   "123",
		Passwd: "456",
	})

	// 查看用户
	var user User
	db.Where("name = ?", 123).First(&user)
	fmt.Println(user)

	// 修改用户
	user.Passwd = "90987"
	db.Save(&user)

	// 删除用户
	db.Delete(&user)
}

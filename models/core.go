package models

import (
	"github.com/jinzhu/gorm"
	"os"
)

//定义db变量，变量名开头是小写，只能同包下访问
var db gorm.DB

func init() {
	//创建data目录，0777代表文件权限
	if err := os.MkdirAll("data", 0777); err != nil {
		panic("failed to connect database," + err.Error())
	}
	//打开数据库，并复杂给db变量，data/data.db数据文件路径
	db, err := gorm.Open("sqlite3", "data/data.db")
	//存在错误，则程序退出，panic是类似于java的RuntimeException错误
	if err != nil {
		panic("failed to connect database")
	}

	//自动同步表结构
	db.AutoMigrate(&User{})

	var count int
	//Model(&User{})查询用户表，Count(&Count)将用户表的数据赋值给count字段。
	if err := db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
		//新增
		db.Create(&User{
			Name:   "admin",
			Email:  "admin@qq.com",
			Pwd:    "123123",
			Avatar: "/static/images/info-img.png",
			Role:   0,
		})
	}
}

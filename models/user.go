package models

import "github.com/jinzhu/gorm"

//定义用户表信息
type User struct {
	gorm.Model
	Name   string `gorm:"unique_index"`
	Email  string `gorm:"unique_index"`
	Avatar string
	Pwd    string
	Role   int `gorm:"default:1"` //0管路员  1正常用户
}

func QueryUserByEmailAndPassword(email, password string) (user User, err error )  {
	return user, db.Where("email = ? and pwd = ?",email, password).Take(&user).Error
}

//根据昵称查用户
func QueryUserByName(name string) (user User, err error)  {
	return user, db.Where("name = ?", name).Take(&user).Error
}

//根据邮箱查用户
func QueryUserByEmail(email string) (user User, err error)  {
	return user, db.Where("email = ?", email).Take(&user).Error
}

//保存用户
func SaveUser(user *User) error  {
	return db.Create(user).Error
}

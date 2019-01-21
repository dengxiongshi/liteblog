package controllers

import (
	"liteblog/models"
	"strings"
	"errors"
)

type UserController struct {
	BaseController
}

//@router /login [post]
func (c *UserController) Login()  {
	//判断邮箱不能为空
	email := c.GetMustString("email", "邮箱不能为空!")
	//判断密码不能为空
	pwd := c.GetMustString("password", "密码不能为空!")
	var (
		user *models.User
		err  error
	)
	if user, err = models.QueryUserByEmailAndPassword(email, pwd); err != nil {
		c.Abort500(syserrors.NewError("邮箱或密码不对", err))
	}
	//将user保存到session
	c.SetSession(SESSION_USER_KEY, user)
	c.JSONK("登录成功", "/")
}

//@router /reg [post]
func (c *UserController) Reg()  {
	name := c.GetMustString("name", "昵称不能为空!")
	email := c.GetMustString("email", "邮箱不能为空!")
	pwd := c.GetMustString("password", "密码不能我空!")
	pwd2 := c.GetMustString("password2", "确认密码不能为空!")
	if strings.Compare(pwd, pwd2) != 0 {
		c.Abort500(errors.New("密码与确认密码必须要一致!"))
	}
	if u, err := models.QueryUserByName(name); err != nil && u.ID != 0 {
		c.Abort500(syserrors.NewError("用户昵称已经存在!", err))
	}
	if u, err := models.QueryUserByEmail(email); err != nil && u.ID != 0 {
		c.Abort500(syserrors.NewError("用户邮箱已经存在!", err))
	}
	//开始保存用户
	if err := models.SaveUser(&models.User{
		Name: name,
		Email: email,
		Pwd: pwd,
		Avatar: "",
		Role: 1,
	}); err != nil {
		c.Abort500(syserrors.NewError("用户注册失败", err))
	}
	c.JSONOK("注册成功", "/user")
}

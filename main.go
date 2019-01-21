package main

import (
	"github.com/astaxie/beego"
	_ "liteblog/models"
	_ "liteblog/routers"
	"strings"
	"encoding/gob"
	"liteblog/models"
)

func main() {
	initTemplate()
	initSession()
	beego.Run()
}

func initTemplate() {
	beego.AddFuncMap("equrl", func(x, y string) bool {
		s1 := strings.Trim(x, "/")
		s2 := strings.Trim(y, "/")
		return strings.Compare(s1, s2) == 0
	})
}

func initSession()  {
	//beego的session序列号是用gob的方式，因此需要将注册models.User
	gob.Register(models.User{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "liteblog-key"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}

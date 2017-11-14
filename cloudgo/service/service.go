package service

import (
	"github.com/go-martini/martini"
)

func NewServer(port string) {
	//创建一个典型的martini实例
	mt := martini.Classic()
	
	//接收对/login的GET方法请求，第二个参数是对该请求的处理方法，获取其参数usr的值并打印
	mt.Get("/login/:usr", func(params martini.Params) string {
		return "user " + params["usr"] + " has logged in.\n"
	})
	
	//运行该服务器
	mt.RunOnAddr(":" + port)
}

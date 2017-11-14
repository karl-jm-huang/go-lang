package service

import (
	"github.com/go-martini/martini"
)

func NewServer(port string) {
	mt := martini.Classic()

	mt.Get("/login/:usr", func(params martini.Params) string {
		return "user " + params["usr"] + " has logged in.\n"
	})

	mt.RunOnAddr(":" + port)
}

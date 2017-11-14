package main

import (
	"cloudgo/service"
	"os"

	flag "github.com/spf13/pflag"
)

//默认端口
const (
	PORT string = "8080"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}
	
	//解析命令行参数看是否传入了监听端口
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	service.NewServer(port)
}

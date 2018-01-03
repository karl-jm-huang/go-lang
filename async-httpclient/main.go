package main

import (
	"async-httpclient/client"
	"fmt"
	"time"
)

func syncTest() {
	start := time.Now()

	x := client.HTTPGet("guangzhou")
	y := client.HTTPGet("wuhan")
	z := client.HTTPGet("shenzhen")
	m := client.HTTPGet("zhuhai")
	n := client.HTTPGet("chongqing")

	fmt.Printf("%s", x+y+z+m+n)
	fmt.Println("\n====== 同步总耗时: ", time.Since(start), " ======")
}

func asyncTest() {
	s := make(chan string)
	start := time.Now()

	go client.HTTPGetAsync("tianjin", s)
	go client.HTTPGetAsync("wuhan", s)
	go client.HTTPGetAsync("shenzhen", s)
	go client.HTTPGetAsync("zhuhai", s)
	go client.HTTPGetAsync("shijiazhuang", s)

	fmt.Printf("%s", <-s+<-s+<-s+<-s+<-s)
	fmt.Println("\n====== 基于消息的异步机制总耗时: ", time.Since(start), " ======")
}

func main() {
	fmt.Printf("\n%s\n", "====== 实现图 6-2 的 Naive Approach ======")
	syncTest()

	fmt.Printf("\n%s\n", "====== 实现图 6-3，为每个 HTTP 搭建基于消息的异步机制 ======")
	asyncTest()
}

package main

import "log"

// 这一组引用将自动把模块注册到module的全局变量里
import (
	_ "github.com/funny/go-project-demo/server1/module/item"
	_ "github.com/funny/go-project-demo/server1/module/player"
)

func main() {
	log.Println("server1 start")
	log.Println("server1 exit")
}

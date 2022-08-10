package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rank_list/controller/rank"
)

const (
	ServerAddr = "127.0.0.1:7700"
)

func main() {
	e := gin.Default()
	//...
	initRoute(e)
}


func initRoute(e *gin.Engine) {
	rank.Route(e)
	if err := e.Run(ServerAddr); err != nil {
		fmt.Println("main.initRoute failed, err=", err)
		return
	}
}
package main

import (
	"go_chat/cache"
	"go_chat/conf"
	"go_chat/router"
)

func main() {
	conf.Init()
	cache.Init()
	r := router.NewRouter()
	_ = r.Run(conf.HttpPort)
}

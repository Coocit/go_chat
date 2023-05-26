package main

import (
	"go_chat/cache"
	"go_chat/conf"
)

func main() {
	conf.Init()
	cache.Init()
}

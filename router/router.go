package router

import (
	"github.com/gin-gonic/gin"
	"go_chat/api"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	/*
		Recovery() 中间件会恢复(recovers) 任何恐慌(panics)
		如果存在恐慌，中间件将会写入500。
		这个中间件还是很必要的，因为你的程序里有些异常情况你没考虑到的时候，程序就退出了
		服务就停止了，所以是必要的
	*/
	r.Use(gin.Recovery(), gin.Logger())
	v1 := r.Group("/")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		v1.POST("user/register", api.UserRegister)
	}
	return r
}

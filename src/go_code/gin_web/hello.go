package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	//定义路由的GET方法及响应处理函数
	engine.GET("/hello", func(ctx *gin.Context) {
		//将发送的信息封装成JSON发送给浏览器
		ctx.JSON(http.StatusOK, gin.H{
			//这是我们定义的数据
			"message": "快速入门",
		})
	})
	engine.Run() //默认在本地8080端口启动服务
}

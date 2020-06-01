package main

import (
	"fmt"
	"hello_gin/app/core"
	"hello_gin/app/global"
	"hello_gin/app/initial"
)

func main() {
	fmt.Println("hello gin")
	core.RunApp()
	// r := gin.Default()
	initial.Mysql()
	initial.DBTables()
	r := initial.Routers()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// 程序结束前关闭数据库链接
	defer global.GVA_DB.Close()
	fmt.Println("listen and serve on 0.0.0.0:8080")
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

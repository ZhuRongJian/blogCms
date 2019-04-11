package main

import (
	"blog-cms/conf"
	"blog-cms/router"
	"fmt"
)

func main() {
	// 初始化各项配置
	conf.SetUp()

	//设置路由
	route := router.SetUpRouter()
	err := route.Run(":8080")

	if err != nil {
		fmt.Println(err)
	}
}

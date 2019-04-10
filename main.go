package main

import (
	"blog-cms/router"
)

func main() {
	route := router.SetUpRouter()
	route.Run(":8080")
}

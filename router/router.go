package router

import (
	"blog-cms/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpRouter() *gin.Engine {

	router := gin.Default()
	// 跨域注册cors
	router.Use(middleware.CORSMiddleware())

	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./static/images/logo.jpg")
	router.LoadHTMLGlob("views/*")
	//后台管理员页面
	admin := router.Group("/admin")
	{
		admin.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})
		admin.GET("/one", func(c *gin.Context) {
			c.HTML(http.StatusOK, "one.html", gin.H{})
		})
		admin.GET("/two", func(c *gin.Context) {
			c.HTML(http.StatusOK, "two.html", gin.H{})
		})
	}

	return router
}

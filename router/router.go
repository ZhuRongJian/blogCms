package router

import (
	"blog-cms/controllers/post"
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
	router.LoadHTMLGlob("views/**/*")
	// 后台管理员页面
	admin := router.Group("/admin")
	{
		// 后台页面路由
		adv := admin.Group("/view")
		{
			adv.GET("/index", func(c *gin.Context) {
				c.HTML(http.StatusOK, "admin/index.html", gin.H{})
			})
			adv.GET("/one", func(c *gin.Context) {
				xx := c.DefaultQuery("page", "0")
				c.HTML(http.StatusOK, "admin/one.html", gin.H{
					"title": xx,
				})
			})
			adv.GET("/two", func(c *gin.Context) {
				c.HTML(http.StatusOK, "admin/two.html", gin.H{})
			})
		}
		// 后台接口路由
		adp := admin.Group("/api")
		{
			adpp := adp.Group("/post")
			{
				adpp.POST("add", post.Add)
			}
		}
	}

	// 前台blog页面
	blog := router.Group("/blog")
	{
		// 前台页面路由
		blogv := blog.Group("/view")
		{
			blogv.GET("/", func(c *gin.Context) {
				c.HTML(http.StatusOK, "blog/index.html", gin.H{})
			})
		}
		// 前台接口路由
	}

	return router
}

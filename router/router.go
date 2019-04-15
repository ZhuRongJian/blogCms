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
			blogv.GET("/about", func(c *gin.Context) {
				c.HTML(http.StatusOK, "blog/about.html", gin.H{})
			})
			blogv.GET("/fengmian", func(c *gin.Context) {
				c.HTML(http.StatusOK, "blog/fengmian.html", gin.H{})
			})
			blogv.GET("/info", func(c *gin.Context) {
				c.HTML(http.StatusOK, "blog/info.html", gin.H{})
			})
			blogv.GET("/info2", func(c *gin.Context) {
				c.HTML(http.StatusOK, "blog/info2.html", gin.H{})
			})
			blogv.GET("/list", func(c *gin.Context) {
				c.HTML(http.StatusOK, "blog/list.html", gin.H{})
			})
			blogv.GET("/time", func(c *gin.Context) {
				c.HTML(http.StatusOK, "blog/time.html", gin.H{})
			})
		}
		// 前台接口路由
	}

	return router
}

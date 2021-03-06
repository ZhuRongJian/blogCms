package post

import (
	"blog-cms/controllers"
	"blog-cms/models"
	"blog-cms/service"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {

}

func Add(c *gin.Context) {
	post := new(models.Post)
	if err := c.ShouldBind(post); err != nil {
		controllers.SendResponse(c, err, nil)
		return
	}
	ser := service.PostService{}
	result, err := ser.Add(post)
	if err != nil {
		controllers.SendResponse(c, err, nil)
	}
	controllers.SendResponse(c, nil, result)
}

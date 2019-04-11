package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"msg"`
	Data      interface{} `json:"data,omitempty"`
	RequestId string      `json:"request_id"`
	PageNum   int         `json:"page_num,omitempty"`
	PageSize  int         `json:"page_size,omitempty"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	msg := "OK"
	if err != nil {
		msg = err.Error()
	}
	c.JSON(http.StatusOK, Response{
		Code:      200,
		Message:   msg,
		Data:      data,
		RequestId: c.GetString("X-Request-Id"),
	})
	c.Abort()
}

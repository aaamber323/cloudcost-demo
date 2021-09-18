package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//ReturnError...
func ReturnError(c *gin.Context, code int, err error) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  err.Error(),
		Data: nil,
	})
	c.Abort()
}

//ReturnJson
func ReturnJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: SUCCESS,
		Msg:  "success",
		Data: data,
	})
	c.Abort() //阻止调用挂起的是处理程序，不会阻止当前的处理程序
}

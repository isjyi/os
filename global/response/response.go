package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func (r Response) result(c *gin.Context) {
	c.JSON(http.StatusOK, r)
}

func success(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func error(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	success(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	success(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	success(SUCCESS, data, "操作成功", c)
}

func OkDetailed(data interface{}, message string, c *gin.Context) {
	success(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	error(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	error(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	error(code, data, message, c)
}

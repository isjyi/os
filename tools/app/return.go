package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/isjyi/os/utils"
)

// 失败数据处理
func Error(c *gin.Context, code int, err error, msg string) {
	var res Response
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		res.Msg = err.Error() // 翻译校验错误提示
	} else {
		res.Data = utils.T(errs)
	}

	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnError(code))
}

// 通常成功数据处理
func OK(c *gin.Context, data interface{}, msg string) {
	var res Response
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}

// 分页数据处理
func PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	var res PageResponse
	res.Data.List = result
	res.Data.Count = count
	res.Data.PageIndex = pageIndex
	res.Data.PageSize = pageSize
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}

// 兼容函数
func Custum(c *gin.Context, data gin.H) {
	c.JSON(http.StatusOK, data)
}

package dict

import (
	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/models"
	"github.com/isjyi/os/server"
	"github.com/isjyi/os/tools"
	"github.com/isjyi/os/tools/app"
	"github.com/isjyi/os/tools/app/msg"
)

// @Summary 通过字典类型id获取字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Param type_id path int true "type_id"
// @Success 200 {object} app.Response{data=[]models.SysDictData} "desc"
// @Router /api/v1/dict/databytype/{type_id} [get]
// @Security Bearer
func GetDictDataByDictTypeId(c *gin.Context) {
	var DictData models.SysDictData

	id, err := tools.StringToUInt64(c.Param("type_id"))
	tools.HasError(err, "查询字段不能为空", -1)

	DictData.DictTypeId = id
	result, err := DictData.Get()

	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result, msg.GetSuccess)
}

// @Summary 字典数据列表
// @Description 获取JSON
// @Tags 字典数据
// @Param status query string false "status"
// @Param dictCode query string false "dictCode"
// @Param dictType query string false "dictType"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/data/list [get]
// @Security Bearer
func GetDictDataList(c *gin.Context) {
	var r server.DictDataQuery

	if err := c.ShouldBind(&r); err != nil {
		app.Error(c, 402, err, "数据验证失败")
		return
	}

	result, count, err := r.GetPage()
	tools.HasError(err, msg.NotFound, -1)

	app.PageOK(c, result, count, r.PageIndex, r.PageSize, msg.GetSuccess)
}

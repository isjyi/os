package dict

import (
	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/models"
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

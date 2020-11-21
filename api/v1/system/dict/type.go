package dict

import (
	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/server"
	"github.com/isjyi/os/tools"
	"github.com/isjyi/os/tools/app"
	"github.com/isjyi/os/tools/app/msg"
)

// @Summary 字典类型列表数据
// @Description 获取JSON
// @Tags 字典数据
// @Param name query string false "字典名称"
// @Param id query string false "字典id"
// @Param type query string false "字典类型"
// @Param page_size query int false "页条数"
// @Param page_index query int false "页码"
// @Success 200 {object} app.PageResponse{data=app.Page{list=[]models.SysDictType}} "desc"
// @Router /api/v1/dict/type [get]
// @Security Bearer
func GetDictTypeList(c *gin.Context) {
	var r server.DictTypeQuery

	if err := c.ShouldBind(&r); err != nil {
		app.Error(c, 402, err, "数据验证失败")
		return
	}
	result, count, err := r.GetPage()
	tools.HasError(err, "", -1)

	app.PageOK(c, result, count, r.PageIndex, r.PageSize, "")
}

// @Summary 添加字典类型
// @Description 获取JSON
// @Tags 字典数据
// @Accept  application/json
// @Product application/json
// @Param data body server.DictTypeCreate true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/dict/type [post]
// @Security Bearer
func InsertDictType(c *gin.Context) {
	var r server.DictTypeCreate

	if err := c.ShouldBindJSON(&r); err != nil {
		app.Error(c, 402, err, "数据验证失败")
		return
	}

	r.CreateBy = tools.GetUserIdUint64(c)

	result, err := r.Create()

	tools.HasError(err, "error", -1)

	app.OK(c, result, msg.CreatedSuccess)
}

// @Summary 通过字典id获取字典类型
// @Description 获取JSON
// @Tags 字典数据
// @Param id path int true "字典id"
// @Success 200 {object} app.Response{data=models.SysDictType}
// @Router /api/v1/dict/type/{id} [get]
// @Security Bearer
func GetDictType(c *gin.Context) {
	var r server.DictTypeQuery

	r.Id, _ = tools.StringToUInt64(c.Param("id"))

	result, err := r.Get()

	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result, msg.GetSuccess)
}

// @Summary 修改字典类型
// @Description 获取JSON
// @Tags 字典类型
// @Accept  application/json
// @Product application/json
// @Param data body server.DictTypeUpdate true "data"
// @Success 200 {string} string	"{"code": 200, "message": "更新成功！"}"
// @Router /api/v1/dict/type [put]
// @Security Bearer
func UpdateDictType(c *gin.Context) {
	var r server.DictTypeUpdate

	if err := c.ShouldBindJSON(&r); err != nil {
		app.Error(c, 402, err, "数据验证失败")
		return
	}

	r.UpdateBy = tools.GetUserIdUint64(c)

	err := r.Update()

	tools.HasError(err, "", -1)

	app.OK(c, "", msg.UpdatedSuccess)
}

// @Summary 删除字典类型
// @Description 删除数据
// @Tags 字典类型
// @Param id path int true "id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/dict/type/{id} [delete]
func DeleteDictType(c *gin.Context) {
	var r server.DictTypeDelect

	r.Ids = tools.IdsStrToIdsIntGroup("id", c)
	err := r.BatchDelete()
	tools.HasError(err, msg.DeletedFail, -1)
	app.OK(c, "", msg.DeletedSuccess)
}

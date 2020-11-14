package server

import (
	"errors"

	"github.com/isjyi/os/global"
	"github.com/isjyi/os/models"
	"github.com/isjyi/os/tools/app/msg"
	"gorm.io/gorm"
)

type DictTypeQuery struct {
	Name string `form:"name" binding:"omitempty"`
	Id   uint64 `form:"id"  binding:"omitempty"`
	Type string `form:"type" binding:"omitempty"`
	PaginationQuery
}

func (d *DictTypeQuery) scopeId(db *gorm.DB) {
	if d.Id != 0 {
		db.Where("id = ?", d.Id)
	}
}

func (d *DictTypeQuery) scopeType(db *gorm.DB) {
	if d.Type != "" {
		db.Where("type = ?", d.Type)
	}
}

func (d *DictTypeQuery) GetPage() ([]models.SysDictType, int, error) {
	var r []models.SysDictType
	db := global.Eloquent.Table(models.SysDictType{}.TableName())
	d.scopeId(db)
	d.scopeType(db)

	var count int64
	res := db.Offset((d.GetPageIndex() - 1) * d.GetPageSize()).Limit(d.GetPageSize()).Find(&r)
	if res.Error != nil {
		return nil, 0, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, 0, errors.New(msg.NotFound)
	}

	res = db.Offset(-1).Limit(-1).Count(&count)
	if res.Error != nil {
		return nil, 0, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, 0, errors.New(msg.NotFound)
	}
	return r, int(count), nil
}

func (d *DictTypeQuery) Get() (models.SysDictType, error) {
	var r models.SysDictType

	if err := global.Eloquent.First(&r, d.Id).Error; err != nil {
		return r, err
	}

	return r, nil
}

type DictTypeCreate struct {
	Name     string `json:"name" binding:"required,max=30"`
	Type     string `json:"type" binding:"required,max=100"`
	Remark   string `json:"remark" binding:"omitempty,max=255"`
	CreateBy uint64 `json:""`
	Status   uint8  `json:"status" binding:"required,oneof=0 1"`
}

func (d *DictTypeCreate) Create() (models.SysDictType, error) {
	var dict = models.SysDictType{Name: d.Name, Type: d.Type, Status: d.Status, CreateBy: d.CreateBy, UpdateBy: d.CreateBy, Remark: d.Remark}
	var num int64

	if err := global.Eloquent.Table(dict.TableName()).Where("name = ? or type = ?", d.Name, d.Type).Count(&num).Error; err != nil {
		return dict, err
	}

	if num > 0 {
		return dict, errors.New("字典名称或者字典类型已经存在！")
	}

	result := global.Eloquent.Create(&dict)

	if result.Error != nil {
		err := result.Error
		return dict, err
	}
	return dict, nil
}

type DictTypeUpdate struct {
	Id uint64 `form:"id"  binding:"required"`
	DictTypeCreate
	UpdateBy uint64 `json:""`
}

func (d *DictTypeUpdate) Update() error {
	var dict models.SysDictType

	if err := global.Eloquent.First(&dict, d.Id).Error; err != nil {
		return err
	}

	if err := global.Eloquent.Model(&dict).Updates(map[string]interface{}{"status": d.Status, "update_by": d.UpdateBy, "remark": d.Remark}).Error; err != nil {
		return err
	}

	return nil
}

type DictTypeDelect struct {
	Ids []int
}

func (d *DictTypeDelect) BatchDelete() error {
	if err := global.Eloquent.Where("id in (?)", d.Ids).Delete(&models.SysDictType{}).Error; err != nil {
		return err
	}

	return nil
}

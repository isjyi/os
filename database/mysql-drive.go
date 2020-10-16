package database

import (
	"fmt"
	"time"

	"github.com/isjyi/os/global"
	"github.com/isjyi/os/tools"
	"github.com/isjyi/os/tools/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
)

type Mysql struct {
}

func (m *Mysql) Setup() {
	var err error
	global.Eloquent, err = m.Open()
	if err != nil {
		global.Logger.Fatal(tools.Red(m.GetDriver()+" connect error :"), zap.Error(err))
	} else {
		global.Logger.Info(tools.Green(m.GetDriver() + " connect success !"))
	}
	if global.Eloquent.Error != nil {
		global.Logger.Fatal(tools.Red(" database error :"), zap.Error(global.Eloquent.Error))
	}
	// 数据表前缀
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return "os_" + defaultTableName
	// }

	global.Eloquent.LogMode(config.OSConfig.Logger.EnabledDB)

	global.Eloquent.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	global.Eloquent.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	global.Eloquent.Callback().Delete().Replace("gorm:delete", deleteCallback)
}

// 打开数据库连接
func (m *Mysql) Open() (db *gorm.DB, err error) {
	return gorm.Open(m.GetDriver(), m.GetConnect())
}

// 获取数据库连接
func (m *Mysql) GetConnect() string {
	return config.OSConfig.Database.Source
}

func (m *Mysql) GetDriver() string {
	return config.OSConfig.Database.Driver
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()

		if createdAtField, ok := scope.FieldByName("CreatedAt"); ok {
			if createdAtField.IsBlank {
				createdAtField.Set(nowTime)
			}
		}

		if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
			if updatedAtField.IsBlank {
				updatedAtField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdatedTime", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedTime")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

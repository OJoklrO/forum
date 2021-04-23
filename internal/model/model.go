package model

import (
	"fmt"
	"github.com/OJoklrO/forum/global"
	"github.com/OJoklrO/forum/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime)
	db, err := gorm.Open(databaseSetting.DBType, s)
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	db.SingularTable(true)

	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

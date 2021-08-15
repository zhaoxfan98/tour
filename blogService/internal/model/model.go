package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zhaoxfan98/blog/global"
	"github.com/zhaoxfan98/blog/pkg/setting"
)

//创建公共model
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

//针对创建DB实例的NewDBEngine方法 同时增加gorm开源库的引入和MySQL驱动库的初始化
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	// db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s/%s?charset=%s&parseTime=%t&loc=Local",
	// 	// databaseSetting.UserName,
	// 	// databaseSetting.Password,
	// 	// databaseSetting.Host,
	// 	// databaseSetting.DBName,
	// 	// databaseSetting.Charset,
	// 	// databaseSetting.ParseTime,
	// 	"root",
	// 	"301421",
	// 	"127.0.0.1:3306",
	// 	"blog_service",
	// 	"utf8mb4",
	// 	true,
	// ))

	db, err := gorm.Open("mysql", "root:301421@tcp(127.0.0.1:3306)/blog_service")
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

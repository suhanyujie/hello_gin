package initial

import (
	"hello_gin/app/config"
	"hello_gin/app/global"
	"hello_gin/app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	oplogging "github.com/op/go-logging"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG config.Server
	GVA_LOG    *oplogging.Logger
)

func init() {
	GVA_DB = global.GVA_DB
	GVA_CONFIG = global.GVA_CONFIG
	GVA_LOG = global.GVA_LOG
}

//初始化数据库并产生数据库全局变量
func Mysql() {
	admin := GVA_CONFIG.Mysql
	if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
		GVA_LOG.Error("DEFAULTDB数据库启动异常", err)
	} else {
		GVA_DB = db
		GVA_DB.DB().SetMaxIdleConns(admin.MaxIdleConns)
		GVA_DB.DB().SetMaxOpenConns(admin.MaxOpenConns)
		GVA_DB.LogMode(admin.LogMode)
	}
}

//注册数据库表专用
func DBTables() {
	db := GVA_DB
	db.AutoMigrate(
		models.GinUser{},
	)
	GVA_LOG.Debug("register table success")
}

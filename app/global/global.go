package global

import (
	"hello_gin/app/config"

	"github.com/jinzhu/gorm"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_LOG    *oplogging.Logger
)

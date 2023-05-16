package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"xiangxiang/jackaroo/config"
)

var (
	G_CONFIG config.Server
	G_VP     *viper.Viper
	G_DB     *gorm.DB
)

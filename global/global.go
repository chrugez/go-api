package global

import (
	"github.com/chrugez/go-api/pkg/logger"
	"github.com/chrugez/go-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb *gorm.DB
)
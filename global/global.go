package global

import (
	"github.com/chrugez/go-api/pkg/logger"
	"github.com/chrugez/go-api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb *redis.Client
	Mdb *gorm.DB
)
package initialize

import (
	"github.com/chrugez/go-api/global"
	"github.com/chrugez/go-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
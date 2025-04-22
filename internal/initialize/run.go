package initialize

import (
	"fmt"

	"github.com/chrugez/go-api/global"
	"go.uber.org/zap"
)

func Run() {
	//load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading config mysql ", m.Username, m.Password)
	InitLogger()
	global.Logger.Info("Config Log ok!", zap.String("Ok", "success"))
	InitMySql()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}
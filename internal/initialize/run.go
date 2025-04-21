package initialize

import (
	"fmt"

	"github.com/chrugez/go-api/global"
)

func Run() {
	//load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading config mysql ", m.Username, m.Password)
	InitLogger()
	InitMySql()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}
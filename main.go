package main

import (
	"niubi-mall/core"
	"niubi-mall/global"
	initialize "niubi-mall/initilize"
)

func main() {
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库

	core.RunWindowsServer()
}

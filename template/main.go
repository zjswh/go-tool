package main

import (
	"TEMPLATE/config"
	"TEMPLATE/initialize"
	"github.com/zjswh/go-tool/nacos"
)

func main() {
	config.SystemSetUp()
	systemConfig := config.GVA_SYSTEM_CONFIG
	nacos.Setup(systemConfig.NacosIp, systemConfig.NacosPort, systemConfig.AppIp, systemConfig.AppPort, systemConfig.ServerName)

	config.SetUp()

	//加载数据库
	initialize.Mysql()

	//加载redis
	initialize.Redis()

	initialize.RunServer()
}

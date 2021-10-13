package main

import (
	"GoWeb_Template_CLD/common"
	"GoWeb_Template_CLD/common/logger"
	"GoWeb_Template_CLD/common/snowflake"
	"GoWeb_Template_CLD/conf"
	"GoWeb_Template_CLD/dao/mysql"
	"GoWeb_Template_CLD/dao/redis"
	"GoWeb_Template_CLD/router"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	// 1.加载配置

	//configFilePath := "conf/config.yaml" // TODO 记住 更改配置文件路径
	configFilePath := "config.yaml"
	if err := conf.Init(configFilePath); err != nil {
		fmt.Printf("init settings failed,err:%v\n", err)
		return
	}

	// 2.初始化日志
	if err := logger.Init(conf.Conf.LogConfig, conf.Conf.Mode); err != nil {
		fmt.Printf("init logger failed,err:%v\n", err)
		return
	}
	defer zap.L().Sync() // 将日志写入日志文件中
	fmt.Println("logger init success...")

	// 3.初始化MySQL
	if err := mysql.Init(conf.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed,err:%v\n", err)
		return
	}
	defer mysql.Close()

	// 4.初始化Redis
	if err := redis.Init(conf.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed,err:%v\n", err)
		return
	}
	defer redis.Close()

	// 5.初始化snowflake雪花算法
	if err := snowflake.Init(conf.Conf.StartTime, conf.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed,err:%v\n", err)
		return
	}
	// 6.初始化gin框架内置的校验器使用的翻译器
	if err := common.InitTrans("zh"); err != nil { // zh:中文；en:英文
		fmt.Printf("init validator trans failed,err:%v\n", err)
		return
	}
	// 7.注册路由
	r := router.Setup(conf.Conf.Mode)

	fmt.Println("================================")
	fmt.Println("port：", conf.Conf.Port)
	fmt.Println("================================")

	// 8.优雅开关机
	common.ElegantShutdown(r)
}

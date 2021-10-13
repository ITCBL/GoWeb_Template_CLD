package redis

import (
	"GoWeb_Template_CLD/conf"
	"fmt"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var client *redis.Client

// Init 初始化连接
func Init(cfg *conf.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
		PoolSize: cfg.PoolSize, // 最大连接数
	})

	_, err = client.Ping().Result()
	return
}

// Close 小技巧：因为var db *sqlx.DB ,db首字母为小写，不对外暴露。通过暴露一个函数完成对应的操作。
func Close() {
	_ = client.Close()
}

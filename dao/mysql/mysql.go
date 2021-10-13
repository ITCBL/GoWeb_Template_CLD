package mysql

import (
	"GoWeb_Template_CLD/conf"
	"fmt"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 定义一个全局对象db
var db *sqlx.DB

// Init 定义一个初始化数据库的函数
func Init(cfg *conf.MySQLConfig) (err error) {
	// DSN:Data Source Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		// viper的两种取值方式

		// viper的方法2：通过viper.getType("")
		//viper.GetString("mysql.user"),
		//viper.GetString("mysql.password"),
		//viper.GetString("mysql.host"),
		//viper.GetInt("mysql.port"),
		//viper.GetString("mysql.dbname"),

		// viper的方法2: 通过定义一个struct对象
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)
	// 不会校验账号密码是否正确
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect mysql failed ", zap.Error(err)) // zap.Error()是zap的结构化日志
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

// Close 小技巧：因为var db *sqlx.DB ,db首字母为小写，不对外暴露。通过暴露一个函数完成对应的操作。
func Close() {
	_ = db.Close()
}

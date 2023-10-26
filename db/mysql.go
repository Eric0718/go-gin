package db

import (
	"go-gin/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var SqlEngine *xorm.Engine

func init() {
	log.Println("mysql conf:", config.MysqlConf)
	var err error
	SqlEngine, err = xorm.NewEngine("mysql", config.MysqlConf.User+":"+config.MysqlConf.PassWord+"@tcp("+config.MysqlConf.Host+")/"+config.MysqlConf.DbName+"?charset=utf8")
	if err != nil {
		log.Fatal("错误=", err)
	}
	log.Println("connect db ok!")

	//在控制台打印出生成的SQL语句
	//SqlEngine.ShowSQL(true)

}

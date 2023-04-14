package config

import (
	"log"

	"github.com/go-ini/ini"
)

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatalf("init Config, fail to parse 'config/config.ini': %v", err)
		return
	}

	mapTo("mysql", MysqlConf)
	mapTo("server", HttpConf)
	log.Println("InitConfig:", MysqlConf, HttpConf)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

package config

type DataBase struct {
	User     string
	PassWord string
	Host     string
	DbName   string
}

var MysqlConf = &DataBase{}

type HttpConfig struct {
	HttpPort string
	RunMode  string
}

var HttpConf = &HttpConfig{}

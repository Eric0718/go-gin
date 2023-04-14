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

type LogConfig struct {
	LogMode  int
	FilePath string
	CutSize  int64
	CutByDay bool
	DayAge   int

	GinPath string
	GinName string
}

var LogConf = &LogConfig{}

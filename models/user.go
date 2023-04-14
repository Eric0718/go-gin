package models

type UserBase struct {
	Id          int64
	Email       string
	Addr        string
	Pwd         string
	NoId        string
	Code        string
	IdFlag      int
	Pid         int
	PidFlag     string
	TeamNum     int
	People      int
	UseProfit   float64 `xorm:"Decimal"`
	TotalProfit float64 `xorm:"Decimal"`
	CreateTime  string
	TotalCost   float64 `xorm:"Decimal"`
	DelFlag     int
	IsScore     int
}

type UserRequest struct {
	Email string `json:"email"`
	Addr  string `json:"addr"`
	Flag  int    `json:"flag"`
}

package api

const (
	SUCCESS = 0
	ERROR   = iota + 1000
	INVALID_PARAMS
)

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "Request parameter error",
}

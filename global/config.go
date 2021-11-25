package global

import logger2 "github.com/modongning/gocommon/logger"

var MyLog logger2.Logger

func init() {
	MyLog = logger2.NewLogger("./goblog.log")
}

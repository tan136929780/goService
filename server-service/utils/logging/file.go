package logging

import (
	"fmt"
)

var (
	LogSavePath       = "./log/"
	LogSaveName       = "common.log"
	AccessLogSaveName = "access.log"
	SlowLogSaveName   = "slow.log"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	return LogSavePath + LogSaveName
}

func getAccessLogFileFullPath() string {
	return LogSavePath + AccessLogSaveName
}

func getSlowLogFileFullPath() string {
	return LogSavePath + SlowLogSaveName
}

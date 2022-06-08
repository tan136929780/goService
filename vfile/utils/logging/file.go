package logging

import (
	"fmt"
	"visionvera/vfile/utils/config"
)

var (
	GraylogSavePath   = config.GetString("log.grayLogPath")
	LogSavePath       = config.GetString("log.logPath")
	LogSaveName       = "go.log"
	AccessLogSaveName = "go.access"
	SlowLogSaveName   = "go.slow"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	return GraylogSavePath + LogSaveName
}

func getAccessLogFileFullPath() string {
	return LogSavePath + AccessLogSaveName
}

func getSlowLogFileFullPath() string {
	return LogSavePath + SlowLogSaveName
}

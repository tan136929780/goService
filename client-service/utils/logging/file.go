package logging

import (
	"fmt"
)

var (
	GraylogSavePath   = "/applogs/"
	LogSavePath       = "/app/logs/"
	LogSaveName       = "go-detail.app"
	AccessLogSaveName = "go-detail.app.access"
	SlowLogSaveName   = "go-detail.app.slow"
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

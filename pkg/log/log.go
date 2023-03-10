package log

import (
	"github.com/wailsapp/wails/v2/pkg/logger"
)

var (
	Logger logger.Logger
)

const (
	logPath = "leviosa.log"
)

func InitLogger() {
	Logger = logger.NewFileLogger(logPath)

}

package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/astaxie/beego/logs"
)

var logger *logs.BeeLogger

func InitLogger() {
	logger = logs.NewLogger(10000)
	defer logger.Flush()

	path, _ := os.Getwd()
	fileName := path + GetConf().String("log::filename")
	_, err := os.Stat(fileName)
	if err != nil {
		err := os.MkdirAll(filepath.Dir(fileName), os.ModePerm)
		if err != nil {
			panic("无法创建日志目录")
		}
	}

	enableCallDepth, _ := GetConf().Bool("log::enable_func_call_depth")
	logger.EnableFuncCallDepth(enableCallDepth)

	logConf := fmt.Sprintf(`{"filename":"%s"}`, fileName)
	logger.SetLogger("file", logConf)
	logger.SetLevel(getLogLevel())
}

func getLogLevel() int {
	env := GetEnv()
	switch env {
	case ENV_DEV:
		return logs.LevelDebug
	case ENV_TEST:
		return logs.LevelDebug
	case ENV_PROD:
		return logs.LevelError
	default:
		return logs.LevelDebug
	}
}

func GetLog() *logs.BeeLogger {
	return logger
}

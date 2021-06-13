//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package conf

import (
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/util"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var SystemDebugFlag = false

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()

	var logLevel int
	switch os.Getenv("LOG_LEVEL") {
	case "debug":
		logLevel = util.LevelDebug
	case "info":
		logLevel = util.LevelInformational
	case "warn":
		logLevel = util.LevelWarning
	case "error":
		logLevel = util.LevelError
	default:
		logLevel = util.LevelInformational
	}

	if os.Getenv("GIN_MODE") == "debug" {
		SystemDebugFlag = true
	}

	adminDefaultPassword := os.Getenv("ADMIN_PASSWD")
	if adminDefaultPassword == "" || len(adminDefaultPassword) < 8 {
		model.AdminDefaultPasswd = "12345678"
	} else {
		model.AdminDefaultPasswd = adminDefaultPassword
	}

	logMaxSize, err1 := strconv.Atoi(os.Getenv("LOG_MAX_SIZE"))

	logMaxAge, err2 := strconv.Atoi(os.Getenv("LOG_MAX_AGE"))

	logMaxBackUp, err3 := strconv.Atoi(os.Getenv("LOG_MAX_BACKUP"))

	logCompress, err4 := strconv.ParseBool(os.Getenv("LOG_COMPRESS"))

	logJSONFormat, err5 := strconv.ParseBool(os.Getenv("LOG_JSON_FORMAT"))

	logShowLines, err6 := strconv.ParseBool(os.Getenv("LOG_SHOW_LINES"))

	logShowInConsole, err7 := strconv.ParseBool(os.Getenv("LOG_SHOW_IN_CONSOLE"))

	// Set default logger config
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil || err7 != nil {
		fmt.Println("Warn: Custom log config missing. Using default config.")
		logMaxSize = 50
		logMaxAge = 30
		logMaxBackUp = 0
		logCompress = false
		logJSONFormat = false
		logShowLines = true
		logShowInConsole = true
	}

	// 设置日志级别
	util.InitLogger(
		os.Getenv("LOG_PATH"),
		logLevel,
		logMaxSize,
		logMaxBackUp,
		logMaxAge,
		logCompress,
		logJSONFormat,
		logShowLines,
		logShowInConsole,
	)

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", zap.Error(err))
	}

	var DBLogLevel logger.LogLevel
	switch os.Getenv("DB_LOG_LEVEL") {
	case "silent":
		DBLogLevel = logger.Silent
	case "warn":
		DBLogLevel = logger.Warn
	case "error":
		DBLogLevel = logger.Error
	case "info":
		DBLogLevel = logger.Info
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"), DBLogLevel)
	//cache.Redis()
}

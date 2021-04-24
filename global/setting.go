package global

import (
	"forum/pkg/logger"
	"forum/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	JWTSetting      *setting.JWTSettings

	Logger *logger.Logger
)

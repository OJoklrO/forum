package global

import (
	"github.com/OJoklrO/forum/pkg/logger"
	"github.com/OJoklrO/forum/pkg/setting"
)

var (
	ServerSetting           *setting.ServerSettings
	AppSetting              *setting.AppSettings
	DatabaseSetting         *setting.DatabaseSettings

	Logger                  *logger.Logger
)
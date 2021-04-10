package global

import (
	"github.com/GriezLiao/griez-go-tour/blog/pkg/logger"
	"github.com/GriezLiao/griez-go-tour/blog/pkg/setting"
)

var (
	ServerSetting *setting.ServerSetting
	AppSetting    *setting.AppSetting
	DataSetting   *setting.DataBaseSetting
	Logger        *logger.Logger
)

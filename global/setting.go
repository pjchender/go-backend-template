package global

import (
	"gorm.io/gorm"
	"github.com/pjchender/go-backend-template/configs"
)

var (
	AppSetting        *configs.App
	HTTPServerSetting *configs.HTTPServer
	GRPCSetting       *configs.GRPC
	DatabaseSetting   *configs.Database
	AuthSetting       *configs.Auth
	GormSetting       *gorm.Config
)

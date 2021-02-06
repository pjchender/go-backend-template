package setting

import (
	"gorm.io/gorm"
	"github.com/pjchender/go-backend-template/configs"
)

func (s *Setting) ReadGormSetting() *gorm.Config {
	return configs.Gorm
}

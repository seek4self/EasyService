package global

import (
    "{{ .Name }}/model/config"

    "gorm.io/gorm"
    "github.com/seek4self/logger"
)

var (
    DB     *gorm.DB
	CONFIG config.Server
    LOG    *logger.Logger 
)
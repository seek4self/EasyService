package global

import (
    "{{ .Name }}/model/config"

    "gorm.io/gorm"
)

var (
    DB     *gorm.DB
	CONFIG config.Server
)
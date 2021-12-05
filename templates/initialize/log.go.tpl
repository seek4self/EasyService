package initialize

import (
    "{{ .Name }}/global"
    "github.com/seek4self/logger"
)

func init() {
    log, err := logger.NewDefault()
    if err != nil {
        panic(err)
    }
    global.LOG = log
}
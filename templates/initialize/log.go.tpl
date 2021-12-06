package initialize

import (
	"log"
	"{{ .Name }}/global"

	"github.com/seek4self/logger"
)

func init() {
	logger, err := logger.New(logger.Config{
		Level:      global.CONFIG.Log.Level,
		Dir:        global.CONFIG.Log.Director,
		Prefix:     global.CONFIG.Log.Prefix,
		Color:      global.CONFIG.Log.Color,
		Stacktrace: global.CONFIG.Log.Stacktrace,
		ShowLine:   global.CONFIG.Log.ShowLine,
		Encoder:    global.CONFIG.Log.Format,
	})
	if err != nil {
		log.Panicln("init logger err", err)
	}

	global.LOG = logger
}
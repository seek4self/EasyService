package initialize

import (
	"fmt"
	"strings"
	"{{ .Name }}/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const defaultConfigFile = "config.yaml"

func init() {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	// 环境变量前缀
	v.SetEnvPrefix("{{upper .Name }}")
	// 根据配置文件生成对应的环境变量，小写默认转换大写，且环境变量优先
	v.AutomaticEnv()
	// 替换配置文件的字段连接符`.`为`_`
	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
}
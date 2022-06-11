package configs

import (
	"bytes"
	_ "embed"
	"github.com/fsnotify/fsnotify"

	"go-web/pkg/env"
	"go-web/pkg/file"

	"github.com/spf13/viper"
	"io"
	"os"
	"path/filepath"
	_ "time"
)

var config = new(Config)

//用于配置环境
type Config struct {
	//用于设置日志级别
	Logger struct {
		Level int `toml:"level"`
	}
}

var (
	//go:embed dev_configs.toml
	devConfigs []byte

	//go:embed test_configs.toml
	testConfigs []byte

	//go:embed pro_configs.toml
	proConfigs []byte
)

//选择对应的环境配置文件
func init() {
	var r io.Reader

	switch env.Active().Value() {
	case "dev":
		r = bytes.NewReader(devConfigs)
	case "test":
		r = bytes.NewReader(testConfigs)
	case "pro":
		r = bytes.NewReader(proConfigs)
	default:
		r = bytes.NewReader(testConfigs)
	}

	viper.SetConfigType("toml")

	if err := viper.ReadConfig(r); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}
	//对配置文件名拼接
	viper.SetConfigName(env.Active().Value() + "_configs")
	//添加到启动路径
	viper.AddConfigPath("./configs")
	//找到文件路径
	configFile := "./configs/" + env.Active().Value() + "_configs.toml"
	_, ok := file.IsExists(configFile)
	//如果不存在则创建对应的文件
	if !ok {
		if err := os.MkdirAll(filepath.Dir(configFile), 0766); err != nil {
			panic(err)
		}

		f, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if err := viper.WriteConfig(); err != nil {
			panic(err)
		}
	}
	//监视配置文件，重新读取配置数据
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

//获取config对象
func Get() Config {
	return *config
}

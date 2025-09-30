package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	globalConfig *Config
	configOnce   sync.Once
)

// Load 加载配置
func Load() *Config {
	configOnce.Do(func() {
		// 获取当前环境 (默认为dev)
		env := os.Getenv("APP_ENV")
		if env == "" {
			env = "dev"
		}

		// 初始化Viper
		v := viper.New()
		v.SetConfigName(env) // 配置文件名称 (不带扩展名)
		v.SetConfigType("yaml")
		v.AddConfigPath("./config/env") // 配置文件路径
		v.AddConfigPath(".")            // 当前目录

		// 读取配置文件
		if err := v.ReadInConfig(); err != nil {
			panic(fmt.Errorf("读取配置文件失败: %w", err))
		}

		// 解析配置到结构体
		var cfg Config
		if err := v.Unmarshal(&cfg); err != nil {
			panic(fmt.Errorf("解析配置文件失败: %w", err))
		}

		// 设置全局配置
		globalConfig = &cfg

		// 监听配置文件变化 (热更新)
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("配置文件已更新:", e.Name)
			if err := v.Unmarshal(&globalConfig); err != nil {
				fmt.Printf("重新加载配置文件失败: %v\n", err)
			}
		})
	})

	return globalConfig
}

// GetConfig 获取全局配置 (线程安全)
func GetConfig() *Config {
	return globalConfig
}

package config

import (
	"time"
)

// AppConfig 应用配置
type AppConfig struct {
	Env        string        `yaml:"env"`         // 环境: dev/test/prod
	Port       string        `yaml:"port"`        // 服务端口
	Timeout    time.Duration `yaml:"timeout"`     // 请求超时时间
	LogLevel   string        `yaml:"log_level"`   // 日志级别
	JWTSecret  string        `yaml:"jwt_secret"`  // JWT密钥
	EnableCORS bool          `yaml:"enable_cors"` // 是否启用CORS
}

// DBConfig 数据库配置
type DBConfig struct {
	Driver          string        `yaml:"driver"`            // 数据库驱动: mysql/postgres
	DSN             string        `yaml:"dsn"`               // 连接字符串
	MaxOpenConns    int           `yaml:"max_open_conns"`    // 最大打开连接数
	MaxIdleConns    int           `yaml:"max_idle_conns"`    // 最大空闲连接数
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"` // 连接最大生命周期
	LogMode         bool          `yaml:"log_mode"`          // 是否开启SQL日志
}

// RedisConfig Redis配置
type RedisConfig struct {
	Addr     string `yaml:"addr"`     // Redis地址
	Password string `yaml:"password"` // Redis密码
	DB       int    `yaml:"db"`       // Redis数据库
}

// Config 全局配置
type Config struct {
	App   AppConfig   `yaml:"app"`
	DB    DBConfig    `yaml:"db"`
	Redis RedisConfig `yaml:"redis"`
}

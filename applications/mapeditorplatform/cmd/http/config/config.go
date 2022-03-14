package config

import (
	"github.com/Fighting2520/kitgo/common/log/logx"
)

type Config struct {
	BaseConfig `yaml:"-,inline"`
	Logs       logx.LogConf `json:"logs" yaml:"Logs"`
	UserMysql  struct {
		Dsn   string `json:"dsn" yaml:"Dsn"`
		Table struct {
			User string `json:"user" yaml:"User"`
		}
	} `json:"userMysql" yaml:"UserMysql"`
	RateLimit struct {
		Limit float64 `json:"limit" yaml:"Limit"`
		Burst int     `json:"burst" yaml:"Burst"`
	} `json:"rateLimit" yaml:"RateLimit"`
}

type (
	LogConfig struct {
		Mode     string `json:"mode" yaml:"Mode"`         // 日志输出模式： 可选值file | console
		Path     string `json:"path" yaml:"Path"`         // 日志保存目录, 当mode = file 生效
		Level    string `json:"level" yaml:"Level"`       // 日志默认级别，默认info, 可选值为 info |error | severe
		Compress bool   `json:"compress" yaml:"Compress"` // 日志是否压缩
		KeepDays int    `json:"keepDays" yaml:"KeepDays"` // 日志保留天数
	}
	BaseConfig struct {
		Host string `json:"host" yaml:"Host"`
		Port int    `json:"port" yaml:"Port"`
	}
)

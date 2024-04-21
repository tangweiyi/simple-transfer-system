package config

import "time"

const (
	DatabaseConfigPath = "db"
)

type DatabaseConfig struct {
	Host           string        `mapstructure:"host"`
	Port           int           `mapstructure:"port"`
	Username       string        `mapstructure:"username"`
	Password       string        `mapstructure:"password"`
	Database       string        `mapstructure:"database"`
	MaxConnections int           `mapstructure:"max-connections"`
	MaxIdleTime    time.Duration `mapstructure:"max-idle-time"`
}

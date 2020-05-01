package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config represents a simple structure containaing all the environment vars.
type Config struct {
	ListenAddr string `envconfig:"listen_addr" default:":8000"`
	RedisAddr  string `envconfig:"redis_addr" default:"redis.service.consul:6379"`
}

// LoadConfig creates a new Config struct and fills it.
func Load() (c *Config, err error) {
	var cfg Config
	err = envconfig.Process("", &cfg)
	return &cfg, err
}

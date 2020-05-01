package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config represents a simple structure containaing all the environment vars.
type Config struct {
	ListenAddr  string `envconfig:"listen_addr" default:":8080"`
	ServiceAddr string `envconfig:"service_addr" default:"service.service.consul"`
}

// LoadConfig creates a new Config struct and fills it.
func Load() (c *Config, err error) {
	var cfg Config
	err = envconfig.Process("", &cfg)
	return &cfg, err
}

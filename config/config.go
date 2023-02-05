package config

import (
	"go-redis/config/cache"
	"os"
	"strconv"
)

type config struct {
}

type Config interface {
	Redis() cache.Redis
	ServiceName() string
	ServicePort() int
	ServiceEnvirontment() string
}

func (c *config) Redis() cache.Redis {
	return cache.InitRedis()
}

func NewConfig() Config {
	return &config{}
}

func (c *config) ServiceName() string {
	return os.Getenv("SERVICE_NAME")
}

func (c *config) ServicePort() int {
	e := os.Getenv("PORT")
	port, _ := strconv.Atoi(e)

	return port
}

func (c *config) ServiceEnvirontment() string {
	return os.Getenv("ENV")
}

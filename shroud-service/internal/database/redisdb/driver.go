package redisdb

import (
	"github.com/redis/go-redis/v9"
)

type config struct {
	Uri      string
	Password string
}

var conf config

func LoadConfig(uri, pass string) {
	conf.Uri = uri
	conf.Password = pass
}

func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     conf.Uri,
		Password: conf.Password,
		DB:       0,
	})
}

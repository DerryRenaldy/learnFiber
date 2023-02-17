package redis

import (
	"github.com/DerryRenaldy/logger/logger"
	"github.com/go-redis/redis/v8"
)

type Connection struct {
	l logger.ILogger
}

type RedisConnection interface {
	RedisConnect()
}

func NewRedisConnection(logger logger.ILogger) *Connection {
	return &Connection{l: logger}
}

func (c Connection) RedisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}

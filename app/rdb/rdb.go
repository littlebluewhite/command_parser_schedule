package rdb

import (
	"command_parser_schedule/util/config"
	"fmt"
	"github.com/redis/go-redis/v9"
	"path/filepath"
	"runtime"
)

var (
	rootPath string
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	rootPath = filepath.Dir(filepath.Dir(filepath.Dir(b)))
}

func NewRedis(yamlName string) *redis.Client {
	redisConfig := config.NewConfig[config.RedisConfig](rootPath, "env", yamlName)
	dsn := fmt.Sprintf("redis://%s:%s@%s:%s/%s",
		redisConfig.User, redisConfig.Password, redisConfig.Host, redisConfig.Port, redisConfig.DB)
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		panic(err)
	}
	return redis.NewClient(opt)
}

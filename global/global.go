package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mmfile/config"
	"mmfile/store"
)

var (
	ServerRedis *redis.Client
	ServerConf  config.ServConf
	ServerViper *viper.Viper
	ServerLog   *zap.Logger
	SqlStore    store.Store
)

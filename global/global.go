package global

import (
	"go.uber.org/zap"
	"mmrights/store"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"mmrights/config"
)

var (
	ServerDb    *gorm.DB
	ServerRedis *redis.Client
	ServerConf  config.ServConf
	ServerViper *viper.Viper
	ServerLog   *zap.Logger
	Store       store.Store
)

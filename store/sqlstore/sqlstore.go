package sqlstore

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SqlStore interface {
	getMaster() *gorm.DB
	getLogger() *zap.Logger
}
package sqlstore

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mmfile/store"
)

type SqlSupplierStores struct {
	rightsRelationAction store.RightsRelationActionStore
}

type SqlSupplier struct {
	stores SqlSupplierStores
	master *gorm.DB
	logger *zap.Logger
}

func (ss *SqlSupplier) InitConnection(masterDb *gorm.DB, logger *zap.Logger) {
	fmt.Println("数据库仓库初始化...")
	ss.master = masterDb
	ss.logger = logger

	ss.stores.rightsRelationAction = NewSqlRightsRelationActionStore(ss)
}

func NewSqlSupplier(defaultDb *gorm.DB, logger *zap.Logger) *SqlSupplier {
	supplier := &SqlSupplier{}
	supplier.InitConnection(defaultDb, logger)

	return supplier
}

func (ss *SqlSupplier) getMaster() *gorm.DB {
	return ss.master
}

func (ss *SqlSupplier) getLogger() *zap.Logger {
	return ss.logger
}

func (ss *SqlSupplier) RightsRelationAction() store.RightsRelationActionStore {
	return ss.stores.rightsRelationAction
}

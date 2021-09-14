package sqlstore

import (
	"fmt"
	"mmrights/store"
)

type  SqlSupplierStores struct {
	rightsRole           store.RightsRoleStore
	rightsRoleGroup      store.RightsRoleGroupStore
	rightsAction         store.RightsActionStore
	rightsActionGroup    store.RightsActionGroupStore
	rightsRelationAction store.RightsRelationActionStore
	rightsUserRole       store.RightsUserRoleStore
	rightsMenuGroup      store.RightsMenuGroupStore
	rightsMenu           store.RightsMenuStore
	rightsRoleMenu       store.RightsRoleMenuStore

}

type SqlSupplier struct {
	stores SqlSupplierStores
}

func NewSqlSupplier() *SqlSupplier {
	supplier := &SqlSupplier{}

	supplier.initConnection()

	supplier.stores.rightsRelationAction = NewSqlRightsRelationActionStore(supplier)
	supplier.stores.rightsUserRole = NewSqlRightsUserRoleStore(supplier)


	return supplier
}

func (ss *SqlSupplier) initConnection() {
	fmt.Println("数据库仓库初始化...")
}


func (ss *SqlSupplier) RightsRole() store.RightsRoleStore {
	return ss.stores.rightsRole
}

func (ss *SqlSupplier) RightsRoleGroup() store.RightsRoleGroupStore {
	return ss.stores.rightsRoleGroup
}

func (ss *SqlSupplier) RightsAction() store.RightsActionStore {
	return ss.stores.rightsAction
}

func (ss *SqlSupplier) RightsActionGroup() store.RightsActionGroupStore {
	return ss.stores.rightsActionGroup
}

func (ss *SqlSupplier) RightsRelationAction() store.RightsRelationActionStore {
	return ss.stores.rightsRelationAction
}

func (ss *SqlSupplier) RightsUserRole() store.RightsUserRoleStore {
	return ss.stores.rightsUserRole
}

func (ss *SqlSupplier) RightsMenuGroup() store.RightsMenuGroupStore {
	return ss.stores.rightsMenuGroup
}

func (ss *SqlSupplier) RightsMenu() store.RightsMenuStore {
	return ss.stores.rightsMenu
}

func (ss *SqlSupplier) RightsRoleMenu() store.RightsRoleMenuStore {
	return ss.stores.rightsRoleMenu
}

package store

import (
	"mmrights/model"
)

type Store interface {
	RightsRole() RightsRoleStore
	RightsRoleGroup() RightsRoleGroupStore
	RightsAction() RightsActionStore
	RightsActionGroup() RightsActionGroupStore
	RightsRelationAction() RightsRelationActionStore
	RightsUserRole() RightsUserRoleStore
	RightsMenuGroup() RightsMenuGroupStore
	RightsMenu() RightsMenuStore
	RightsRoleMenu() RightsRoleMenuStore
}

type RightsRoleStore interface {
}

type RightsRoleGroupStore interface {
}

type RightsActionStore interface {
}

type RightsActionGroupStore interface {
}

type RightsRelationActionStore interface {
	GetActionPackListByUserIds(userId []string) ([]*model.RightsAction, *model.AppError)
}

type RightsUserRoleStore interface {
	GetRoleListByUserId(userId string) ([]*model.RightsRole, *model.AppError)
}

type RightsMenuGroupStore interface {
}

type RightsMenuStore interface {
}
type RightsRoleMenuStore interface {
}
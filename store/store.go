package store

import (
	"mmfile/model"
)

type Store interface {
	RightsRelationAction() RightsRelationActionStore
}

type RightsRelationActionStore interface {
	GetActionPackListByUserIds(userId []string) ([]*model.RightsAction, *model.AppError)
}
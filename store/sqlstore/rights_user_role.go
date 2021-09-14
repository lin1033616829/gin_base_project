// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package sqlstore

import (
	"go.uber.org/zap"
	"mmrights/global"
	"mmrights/model"
	"mmrights/store"
)

type SqlRightsUserRoleStore struct {
	SqlStore
}

func NewSqlRightsUserRoleStore(sqlStore SqlStore) store.RightsUserRoleStore {
	s := &SqlRightsUserRoleStore{
		SqlStore: sqlStore,
	}

	return s
}

func (s SqlRightsUserRoleStore) GetRoleListByUserId(userId string) ([]*model.RightsRole, *model.AppError) {
	rightsRoleList := make([]*model.RightsRole, 0)

	err := global.ServerDb.Table("RightsUserRoles A").
		Select("A.*").
		Joins("RightsRoles B on A.RoleId=B.Id").
		Where("A.UserId = ?", userId).
		Where("B.DeleteAt = ?", 0).
		Find(&rightsRoleList).Error
	if err != nil {
		global.ServerLog.Error("SqlRightsUserRoleStore.GetRoleListByUserId sql err", zap.Error(err))
		return nil, model.NewException(model.ErrDbExecError)
	}

	return rightsRoleList, nil
}

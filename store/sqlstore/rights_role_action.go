// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package sqlstore

import (
	"fmt"
	"go.uber.org/zap"
	"mmfile/model"
	"mmfile/store"
)

type SqlRightsRelationActionStore struct {
	SqlStore
}

func NewSqlRightsRelationActionStore(sqlStore SqlStore) store.RightsRelationActionStore {
	s := &SqlRightsRelationActionStore{
		SqlStore: sqlStore,
	}

	return s
}

// GetActionPackListByUserIds 获取权限关联列表
func (s SqlRightsRelationActionStore) GetActionPackListByUserIds(userId []string) ([]*model.RightsAction, *model.AppError) {
	s.getLogger().Debug(fmt.Sprintf("GetActionPackListByUserIds 传入参数 userIds %v", userId))
	var actionList = make([]*model.RightsAction, 0)

	selectFields := "B.*"
	userSql := fmt.Sprintf("select %v from %v A left join %v B on A.ActionId=B.Id where A.UserId='%v' ",
		selectFields, model.RightsRelationActionTableName, model.RightsActionTableName, userId)
	roleSql := fmt.Sprintf("select %v from %v A left join %v B on A.ActionId=B.Id where A.UserId='%v' or A.RoleId in(select RoleId from RightsUserRoles where UserId='%v')",
		selectFields, model.RightsRelationActionTableName, model.RightsActionTableName, userId, userId)
	querySql := fmt.Sprintf("%v union %v", userSql, roleSql)

	s.getLogger().Debug(fmt.Sprintf("GetActionPackListByUserIds 最终组合sql---------- %v", querySql))

	err := s.getMaster().Raw(querySql).Scan(actionList).Error
	if err != nil {
		s.getLogger().Error("SqlRightsUserRoleStore.GetRoleListByUserId sql err", zap.Error(err))
		return nil, model.NewException(model.ErrDbExecError)
	}

	return actionList, nil

}
package model

const RightsActionGroupTableName = "RightsActionGroups"

type RightsActionGroup struct {
	Id       string `json:"id"`                                    //ID
	Name     string `json:"name" validate:"required,min=2,max=50"` //角色名称
	Weight   int    `json:"weight"`                                //排序 倒叙
	CreateAt int64  `json:"create_at"`                             //创建时间
	DeleteAt int64  `json:"delete_at"`                             //删除时间
}
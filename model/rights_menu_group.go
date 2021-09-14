package model

const (
	RightsMenuGroupTableName = "RightsMenuGroups"
)

type RightsMenuGroup struct {
	Id       string `json:"id"`                                    //ID
	Name     string `json:"name" validate:"required,min=2,max=50"` //分组名称
	Value    string `json:"value"`                                 //跳转值
	Weight   int    `json:"weight"`                                //排序 倒叙
	CreateAt int64  `json:"create_at"`                             //创建时间
	DeleteAt int64  `json:"delete_at"`                             //删除时间
}

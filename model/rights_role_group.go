package model

const (
	RightsRoleGroupTableName = "RightsRoleGroups"
	RightsRoleGroupBasicYes  = 1 //只读
	RightsRoleGroupBasicNo   = 2 //可读可写

	RoleGroupBasicDefaultName = "默认组"
	RoleGroupBasicJobName     = "职务组"
	RoleGroupBasicCustomName  = "自定义组"
)

type RightsRoleGroup struct {
	Id       string `json:"id"`        //ID
	TeamId   string `json:"team_id"`   //TeamID
	Name     string `json:"name"`      //角色组名称
	IsBasic  int8   `json:"is_basic"`  //是否是基础默认 默认为1 基础类型  2不是基础类型 基础默认的为自动插入，禁止修改
	CreateAt int64  `json:"create_at"` //创建时间
	DeleteAt int64  `json:"delete_at"` //删除时间
}
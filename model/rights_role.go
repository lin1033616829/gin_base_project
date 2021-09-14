package model

const (
	RightsRoleTableName = "RightsRoles"

	RoleBasicYes = 1 //基础类型
	RoleBasicNo  = 2 //不是基础类型

	RoleBasicTypeSystem      = 1 //系统管理员
	RoleBasicTypeTeamManager = 2 //团队管理员
	RoleBasicTypeSubManager  = 3 //子管理员

	RoleBasicTypeSystemName      = "系统管理员"
	RoleBasicTypeTeamManagerName = "团队管理员"
	RoleBasicTypeSubManagerName  = "子管理员"
)

// MutexBasicType 这几种基础类型是互斥的
var MutexBasicType = []int{RoleBasicTypeSystem, RoleBasicTypeTeamManager, RoleBasicTypeSubManager}

type RightsRole struct {
	Id        string `json:"id"`         //ID
	GroupId   string `json:"group_id"`   //角色分组ID
	TeamId    string `json:"team_id"`    //TeamID
	Name      string `json:"name"`       //角色名称
	IsBasic   int8   `json:"is_basic"`   //是否是基础默认 默认为1 基础类型  2不是基础类型 基础默认的为自动插入，禁止修改
	BasicType int8   `json:"basic_type"` //基础类型 1 系统管理员 2团队管理员 3子管理员
	CreateAt  int64  `json:"create_at"`  //创建时间
	DeleteAt  int64  `json:"delete_at"`  //删除时间
}
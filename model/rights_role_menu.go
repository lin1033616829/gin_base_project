package model

const (
	RightsRoleMenuTableName = "RightsRoleMenus"
)

type RightsRoleMenu struct {
	Id     string `json:"id"`      //ID
	RoleId string `json:"role_id"` //角色ID
	MenuId string `json:"menu_id"` //菜单ID
	Weight int    `json:"weight"`  //权重
}
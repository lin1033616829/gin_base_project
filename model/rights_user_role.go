package model

const (
	RightsUserRoleTableName = "RightsUserRoles"
)

type RightsUserRole struct {
	Id       string `json:"id"`               //ID
	UserId   string `json:"user_id"`          //用户ID
	RoleId   string `json:"role_id"`          //角色ID
	DepLimit string `json:"department_limit"` //部门限制
	CreateAt int64  `json:"create_at"`        //创建时间
}
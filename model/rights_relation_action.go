package model

const (
	RightsRelationActionTableName = "RightsRelationActions"
)

type RightsRelationAction struct {
	Id       string `json:"id"` //ID
	UserId   string `json:"user_id"`
	ActionId string `json:"action_id"` //行为ID
	RoleId   string `json:"role_id"`   //角色ID
}
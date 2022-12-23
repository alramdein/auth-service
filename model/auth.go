package model

import "context"

type Role string
type Resource string
type Action string

const (
	RoleAdmin  Role = "Admin"
	RoleMember Role = "Member"
)

const (
	UserResource Resource = "UserResource"
)

const (
	Create Action = "Create"
	Update Action = "Update"
	Delete Action = "Delete"
	View   Action = "View"
)

var RoleMap = map[int64]Role{
	1: RoleAdmin,
	2: RoleMember,
}

func GetRole(role_id int64) Role {
	r, err := RoleMap[role_id]
	if err != true {
		return RoleMember
	}
	return r
}

type AuthUsecase interface {
	HasAccess(ctx context.Context, roleID int64, resource Resource, action Action) bool
}

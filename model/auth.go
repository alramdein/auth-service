package model

import (
	"context"

	"github.com/alramdein/auth-service/pb"
)

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

var ResourceMap = map[pb.Resource]Resource{
	pb.Resource_ResourceUser: UserResource,
}

var ActionMap = map[pb.Action]Action{
	pb.Action_ActionCreate: Create,
	pb.Action_ActionUpdate: Update,
	pb.Action_ActionDelete: Delete,
	pb.Action_ActionView:   View,
}

func GetRole(role_id int64) Role {
	r, err := RoleMap[role_id]
	if err != true {
		return RoleMember
	}
	return r
}

func GetResoruce(resource pb.Resource) Resource {
	r, err := ResourceMap[resource]
	if err != true {
		return UserResource
	}
	return r
}

func GetAction(action pb.Action) Action {
	a, err := ActionMap[action]
	if err != true {
		return View
	}
	return a
}

type AuthUsecase interface {
	HasAccess(ctx context.Context, roleID int64, resource Resource, action Action) bool
}

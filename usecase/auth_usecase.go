package usecase

import (
	"context"

	"github.com/alramdein/auth-service/model"
)

type authUsecase struct{}

func NewAuthUsecase() model.AuthUsecase {
	return &authUsecase{}
}

func (u *authUsecase) HasAccess(ctx context.Context, roleID int64, resource model.Resource, action model.Action) bool {
	// TODO: add database regarding the role, resource, action
	if action == model.View {
		return true
	}

	if model.GetRole(roleID) != model.RoleAdmin {
		return false
	}

	return true
}

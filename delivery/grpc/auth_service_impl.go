package grpc

import (
	"context"

	"github.com/alramdein/auth-service/model"
	"github.com/alramdein/auth-service/pb"
)

func (s *Service) HasAccess(ctx context.Context, req *pb.HassAccessRequest) *pb.HasAccessResponse {
	hasAccess := s.authUsecase.HasAccess(
		ctx,
		req.GetRoleId(),
		model.GetResoruce(req.GetResource()),
		model.GetAction(req.GetAction()),
	)
	if hasAccess != true {
		return &pb.HasAccessResponse{HasAccess: false}
	}

	return &pb.HasAccessResponse{HasAccess: true}
}

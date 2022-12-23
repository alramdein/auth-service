package grpc

import (
	"context"

	"github.com/alramdein/auth-service/model"
	"github.com/alramdein/auth-service/pb"
)

func (s *Service) HasAccess(ctx context.Context, req *pb.HassAccessRequest) (res *pb.HasAccessResponse, err error) {
	hasAccess := s.authUsecase.HasAccess(ctx, req.GetRoleId(), model.Resource(req.GetResource()), model.Action(req.GetAction()))
	if hasAccess != true {
		res = &pb.HasAccessResponse{HasAccess: false}
		return
	}

	res = &pb.HasAccessResponse{HasAccess: true}
	return
}

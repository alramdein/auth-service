package grpc

import (
	"github.com/alramdein/auth-service/model"
	"github.com/alramdein/auth-service/pb"
)

type Service struct {
	pb.UnimplementedAuthServiceServer
	authUsecase model.AuthUsecase
}

func NewUserService() *Service {
	return new(Service)
}

func (s *Service) RegisterAuthUsecase(uc model.AuthUsecase) {
	s.authUsecase = uc
}

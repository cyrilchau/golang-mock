package grpc

import (
	"context"

	userUsecase "myapp/internal/user/usecase"
	"myapp/pkg/otel/zerolog"
	pb "myapp/proto/pb"
)

type (
	UserService struct {
		pb.UnimplementedUserServiceServer
		uc  userUsecase.Usecase
		log *zerolog.Logger
	}

	ServiceDeps struct {
		UserUsecaseI userUsecase.Usecase
	}
)

func NewUserService(deps ServiceDeps, log *zerolog.Logger) *UserService {
	return &UserService{
		uc:  deps.UserUsecaseI,
		log: log,
	}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	resp, _, err := s.uc.Detail(ctx, int(req.Id))

	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{Id: int32(resp.UserID), Name: resp.Fullname}, nil
}

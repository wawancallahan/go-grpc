package protoserver

import (
	"context"

	"github.com/wawancallahan/go-grpc/pb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (service *UserServer) Register(_ context.Context, request *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	response := &pb.RegisterUserResponse{
		User: &pb.User{
			Username:  "TES",
			FullName:  "TES",
			Email:     "TES",
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
		},
	}
	return response, nil
}
func (service *UserServer) List(_ context.Context, request *emptypb.Empty) (*pb.ListUserResponse, error) {
	list := make([]*pb.User, 0)
	list = append(list, &pb.User{
		Username:  "TES",
		FullName:  "TES",
		Email:     "TES",
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	})

	response := &pb.ListUserResponse{
		List: list,
	}

	return response, nil
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

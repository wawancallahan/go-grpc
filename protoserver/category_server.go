package protoserver

import (
	"context"

	"github.com/wawancallahan/go-grpc/pb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CategoryServer struct {
	pb.UnimplementedCategoryServiceServer
}

func (service *CategoryServer) Register(_ context.Context, request *pb.RegisterCategoryRequest) (*pb.RegisterCategoryResponse, error) {
	response := &pb.RegisterCategoryResponse{
		Category: &pb.Category{
			Name:      request.GetName(),
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
		},
	}
	return response, nil
}
func (service *CategoryServer) List(_ context.Context, request *emptypb.Empty) (*pb.ListCategoryResponse, error) {
	list := make([]*pb.Category, 0)
	list = append(list, &pb.Category{
		Name:      "TEST",
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	})

	response := &pb.ListCategoryResponse{
		List: list,
	}

	return response, nil
}

func NewCategoryServer() *CategoryServer {
	return &CategoryServer{}
}

package protoclient

import (
	"context"
	"encoding/json"
	"log"

	"github.com/wawancallahan/go-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	ctx = context.Background()
)

type ClientGrpc struct {
	userClient     pb.UserServiceClient
	categoryClient pb.CategoryServiceClient
}

func serveGrpcClient() *ClientGrpc {
	connection, err := grpc.Dial(":3300", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Could not connect client", err)
	}

	return &ClientGrpc{
		userClient:     pb.NewUserServiceClient(connection),
		categoryClient: pb.NewCategoryServiceClient(connection),
	}
}

func main() {
	clientGrpc := serveGrpcClient()

	// response, err := clientGrpc.userClient.List(ctx, new(emptypb.Empty))

	// response, err := clientGrpc.userClient.Register(ctx, &pb.RegisterUserRequest{
	// 	Username: "A",
	// 	FullName: "B",
	// 	Email:    "C",
	// })

	// response, err := clientGrpc.categoryClient.Register(ctx, &pb.RegisterCategoryRequest{
	// 	Name: "D",
	// })

	response, err := clientGrpc.categoryClient.List(ctx, new(emptypb.Empty))

	if err != nil {
		log.Fatal(err.Error())
	}

	listString, _ := json.Marshal(response.GetList())

	log.Println(listString)
}

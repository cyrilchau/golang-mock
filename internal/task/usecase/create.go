package usecase

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"myapp/internal/task/dtos"
	"myapp/internal/task/entity"
	pb "myapp/proto/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:8080", "the address to connect to")
)

func (uc *usecase) CreateNewTask(ctx context.Context, payload dtos.CreateTaskRequest) (result entity.Task, httpCode int, err error) {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	res, err := client.GetUser(ctx, &pb.UserRequest{Id: int32(payload.UserID)})
	if err != nil {
		uc.log.Z().Err(err).Msg("[usecase]client.GetUser")
		return result, http.StatusInternalServerError, err
	}

	if res.Id == 0 {
		err = fmt.Errorf("pb.GetUser: %v", "No user found")
		uc.log.Z().Err(err).Msg("[usecase]client.GetUser")
		return result, http.StatusInternalServerError, err
	}

	result, err = uc.repo.CreateNewTask(ctx, dtos.NewCreateTask(payload, uc.cfg))
	if err != nil {
		uc.log.Z().Err(err).Msg("[usecase]Create.CreateNewTask")

		return result, http.StatusInternalServerError, err
	}

	return result, http.StatusOK, nil
}

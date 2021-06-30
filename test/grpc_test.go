package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	uspb "github.com/muka-id/service-user/applicationservice/client/userclient/generatedclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var client uspb.UserServiceClient
var userId string

const TARGET string = "localhost:5052"

func init() {
	conn, err := grpc.Dial(TARGET, grpc.WithInsecure())
	if err != nil {
		fmt.Errorf("Failed to dial server: %v", err)
	}

	client = uspb.NewUserServiceClient(conn)

}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	user := uspb.User{
		Name: gofakeit.Name(),
		Role: "Admin",
	}

	// Call the gRPC Service
	res, err := client.RegisterUser(ctx, &uspb.StoreUserRequest{User: &user})

	userId = res.User.Id

	if err != nil {
		switch status.Convert(err).Code() {
		case codes.AlreadyExists:
			t.Error("User already exist")
		default:
			t.Error(err)
		}
	}
}

func TestFindUser(t *testing.T) {
	// Find created user

	ctx := context.Background()

	_, err := client.GetUserById(ctx, &uspb.GetUserByIdRequest{Id: userId})

	statusResult := status.Convert(err)

	// Check for errors
	if err != nil {
		switch statusResult.Code() {
		case codes.NotFound:
			t.Error(statusResult.Message())
		case codes.Internal:
			t.Error(statusResult.Message())
		default:
			t.Error(err)
		}
	}
}
func TestUpdateUser(t *testing.T) {
	ctx := context.Background()

	updateData := &uspb.UpdateUserByIdRequest{
		Id: userId,
		User: &uspb.User{
			Id:   userId,
			Name: "UPDATED",
		},
	}
	_, err := client.UpdateUserById(ctx, updateData)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteUser(t *testing.T) {

}

package controller

import (
	"context"

	"github.com/muka-id/service-user/app/container/containerhelper"
	"github.com/muka-id/service-user/app/logger"
	"github.com/muka-id/service-user/applicationservice/client/userclient"
	uspb "github.com/muka-id/service-user/applicationservice/client/userclient/generatedclient"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (userservice *UserService) UpdateUserById(ctx context.Context, request *uspb.UpdateUserByIdRequest) (*uspb.UpdateSingleUserResponse, error) {
	// Get Update Use Case
	usecase, err := containerhelper.GetUpdateUseCase(userservice.Container)
	if err != nil {
		logger.SugarLog.Error("Get Usecase error", err)
		return nil, status.Error(codes.Internal, "Get usecase error")
	}

	// Convert user request to user model
	user, err := userclient.GrpcToUser(request.User)
	if err != nil {
		logger.SugarLog.Error("Get Usecase error", err)
		return nil, status.Error(codes.Internal, "Error convert GRPC request to User model")
	}

	logger.SugarLog.Debug("Update User called for ", user.Id)
	// Run the usecase method
	result, err := usecase.UpdateUserById(user.Id, user)
	if err != nil {
		return nil, err
	}

	// Convert the User model back to GRPC User
	userGrpc, err := userclient.UserToGrpc(result)
	if err != nil {
		logger.SugarLog.Errorf("%+v\n", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	logger.SugarLog.Debug("User updated: ", user)

	// Return response to client
	return &uspb.UpdateSingleUserResponse{
		User: userGrpc,
	}, nil

}

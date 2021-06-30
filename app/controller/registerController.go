package controller

import (
	"context"

	"github.com/muka-id/service-user/app/container/containerhelper"
	"github.com/muka-id/service-user/app/logger"
	"github.com/muka-id/service-user/applicationservice/client/userclient"
	uspb "github.com/muka-id/service-user/applicationservice/client/userclient/generatedclient"
	"github.com/muka-id/service-user/domain/usecase/registration"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (userservice *UserService) RegisterUser(ctx context.Context, request *uspb.StoreUserRequest) (*uspb.StoreUserResponse, error) {
	logger.SugarLog.Debug("RegisterUser called")

	// Get registration Usecase
	ruci, err := containerhelper.GetRegistrationUseCase(userservice.Container)
	if err != nil {
		logger.SugarLog.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}

	// Convert GRPC request to User model
	user, err := userclient.GrpcToUser(request.User)
	if err != nil {
		logger.SugarLog.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	logger.SugarLog.Debug("Register User:", user)

	// Execute register user
	resultUser, err := ruci.RegisterUser(user)
	if err != nil {
		logger.SugarLog.Errorf("%+v\n", err)
		if errors.Is(err, registration.ErrUserExist) {
			st := status.New(codes.AlreadyExists, "User Already Exist")
			return nil, st.Err()
		}
		return nil, err
	}

	// Get response and return back
	response, err := userclient.UserToGrpc(user)
	if err != nil {
		logger.SugarLog.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	logger.SugarLog.Debug("Registered User:", resultUser)

	return &uspb.StoreUserResponse{User: response}, nil
}

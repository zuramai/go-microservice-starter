package controller

import (
	"context"

	"github.com/muka-id/service-user/app/container/containerhelper"
	"github.com/muka-id/service-user/app/logger"
	"github.com/muka-id/service-user/applicationservice/client/userclient"
	uspb "github.com/muka-id/service-user/applicationservice/client/userclient/generatedclient"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (userservice *UserService) GetUserById(ctx context.Context, request *uspb.GetUserByIdRequest) (*uspb.User, error) {
	logger.SugarLog.Debug("GetUserById called")

	listuserUC, err := containerhelper.GetUserUseCase(userservice.Container)
	if err != nil {
		return nil, status.Error(codes.Internal, "Get User Use Case Error")
	}

	user, err := listuserUC.GetUserById(request.Id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			st := status.New(codes.NotFound, "User Not Found")
			st.WithDetails(&errdetails.BadRequest{})
			return nil, st.Err()
		} else {
			logger.SugarLog.Error("GetUserById Error", err)
			return nil, status.Error(codes.Internal, "GetUserById Error")
		}
	}
	// Convert to GRPC Response
	response, err := userclient.UserToGrpc(user)
	if err != nil {
		logger.SugarLog.Error("Error converting user to grpc ", err)
		return nil, errors.Wrap(err, "")
	}

	return response, nil

}

func (userservice *UserService) GetUserList(ctx context.Context, request *uspb.GetUserListRequest) (*uspb.GetUserListResponse, error) {
	return nil, nil
}

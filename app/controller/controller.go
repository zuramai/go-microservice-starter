package controller

import (
	"context"

	"github.com/muka-id/service-user/app/container/servicecontainer"
	uspb "github.com/muka-id/service-user/applicationservice/client/userclient/generatedclient"
)

type UserService struct {
	Container *servicecontainer.ServiceContainer
	uspb.UnimplementedUserServiceServer
}

func (userservice *UserService) mustEmbedUnimplementedUserServiceServer(ctx context.Context, request *uspb.GetUserListRequest) (*uspb.GetUserListResponse, error) {
	return nil, nil
}

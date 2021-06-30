package controller

import (
	"context"

	uspb "github.com/muka-id/service-user/applicationservice/client/userclient/generatedclient"
)

func (userservice *UserService) DeleteUserById(ctx context.Context, request *uspb.DeleteUserByIdRequest) (*uspb.DeleteUserResponse, error) {
	return nil, nil
}

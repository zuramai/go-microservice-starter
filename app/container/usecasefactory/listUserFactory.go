package usecasefactory

import (
	"github.com/muka-id/service-user/app/config"
	"github.com/muka-id/service-user/app/container"
	listuser "github.com/muka-id/service-user/domain/usecase/getuser"
	"github.com/pkg/errors"
)

type ListUserFactory struct{}

func (luf *ListUserFactory) Build(c container.Container, appConfig *config.AppConfig, key string) (UseCaseInterface, error) {
	usecaseconfig := appConfig.UseCaseConfig.ListUser

	udi, err := buildUserData(c, &usecaseconfig.UserDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	listUserUC := listuser.ListUserUseCase{UserDataInterface: udi}

	return &listUserUC, nil
}

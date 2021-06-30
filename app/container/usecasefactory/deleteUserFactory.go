package usecasefactory

import (
	"github.com/muka-id/service-user/app/config"
	"github.com/muka-id/service-user/app/container"
	"github.com/muka-id/service-user/app/logger"
	"github.com/muka-id/service-user/domain/usecase/deleteuser"
	"github.com/pkg/errors"
)

type DeleteUserFactory struct{}

func (duf *DeleteUserFactory) Build(c container.Container, appConfig *config.AppConfig, key string) (UseCaseInterface, error) {
	deleteUserConfig := appConfig.UseCaseConfig.DeleteUser
	udi, err := buildUserData(c, &deleteUserConfig.UserDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	deleteUseCase := deleteuser.DeleteUserUseCase{UserDataInterface: udi}
	logger.SugarLog.Debug("Delete usecase: ", deleteUseCase)
	return &deleteUseCase, nil
}

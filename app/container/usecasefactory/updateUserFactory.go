package usecasefactory

import (
	"github.com/muka-id/service-user/app/config"
	"github.com/muka-id/service-user/app/container"
	"github.com/muka-id/service-user/app/logger"
	"github.com/muka-id/service-user/domain/usecase/updateuser"
	"github.com/pkg/errors"
)

type UpdateUserFactory struct{}

func (uuf *UpdateUserFactory) Build(c container.Container, appConfig *config.AppConfig, key string) (UseCaseInterface, error) {
	updateUserConfig := appConfig.UseCaseConfig.UpdateUser
	udi, err := buildUserData(c, &updateUserConfig.UserDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	updateUseCase := updateuser.UpdateUserUseCase{UserDataInterface: udi}
	logger.SugarLog.Debug("Registration usecase: ", updateUseCase)
	return &updateUseCase, nil
}

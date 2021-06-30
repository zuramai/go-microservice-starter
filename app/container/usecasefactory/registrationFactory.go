package usecasefactory

import (
	"github.com/muka-id/service-user/app/config"
	"github.com/muka-id/service-user/app/container"
	"github.com/muka-id/service-user/app/logger"
	"github.com/muka-id/service-user/domain/usecase/registration"
	"github.com/pkg/errors"
)

type RegistrationFactory struct {
}

func (rf *RegistrationFactory) Build(c container.Container, appConfig *config.AppConfig, key string) (UseCaseInterface, error) {
	registrationConfig := appConfig.UseCaseConfig.Registration
	udi, err := buildUserData(c, &registrationConfig.UserDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	registrationUseCase := registration.RegistrationUseCase{UserDataInterface: udi}
	logger.SugarLog.Debug("Registration usecase: ", registrationUseCase)
	return &registrationUseCase, nil
}

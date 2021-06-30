// containerhelper implements factory pattern to create concrete usecase object
package containerhelper

import (
	"reflect"

	"github.com/muka-id/service-user/app/config"
	"github.com/muka-id/service-user/app/container"
	"github.com/muka-id/service-user/app/logger"
	"github.com/muka-id/service-user/domain/usecase"
	"github.com/pkg/errors"
)

func GetRegistrationUseCase(c container.Container) (usecase.RegistrationUseCaseInterface, error) {
	key := config.REGISTRATION
	useCase, err := c.BuildUseCase(key)

	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	logger.SugarLog.Debug("usecase type: ", reflect.TypeOf(useCase))
	return useCase.(usecase.RegistrationUseCaseInterface), nil
}

func GetUserUseCase(c container.Container) (usecase.GetUserUseCaseInterface, error) {
	logger.SugarLog.Debug("Get user use case")
	key := config.LIST_USER

	useCase, err := c.BuildUseCase(key)

	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	logger.SugarLog.Debug("usecase type: ", reflect.TypeOf(useCase))
	return useCase.(usecase.GetUserUseCaseInterface), nil
}

func GetDeleteUseCase(c container.Container) (usecase.DeleteUserUseCaseInterface, error) {
	logger.SugarLog.Debug("Delete user use case")
	key := config.DELETE_USER

	useCase, err := c.BuildUseCase(key)

	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	logger.SugarLog.Debug("usecase type: ", reflect.TypeOf(useCase))
	return useCase.(usecase.DeleteUserUseCaseInterface), nil
}

func GetUpdateUseCase(c container.Container) (usecase.UpdateUserUseCaseInterface, error) {
	logger.SugarLog.Debug("Update user use case")
	key := config.UPDATE_USER

	useCase, err := c.BuildUseCase(key)

	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	logger.SugarLog.Debug("usecase type: ", reflect.TypeOf(useCase))
	return useCase.(usecase.UpdateUserUseCaseInterface), nil
}

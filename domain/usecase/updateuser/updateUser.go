package updateuser

import (
	"github.com/muka-id/service-user/app/logger"
	"github.com/muka-id/service-user/applicationservice/dataservice"
	"github.com/muka-id/service-user/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateUserUseCase struct {
	UserDataInterface dataservice.UserDataInterface
}

func (updateuu *UpdateUserUseCase) UpdateUserById(id primitive.ObjectID, user *model.User) (*model.User, error) {
	logger.SugarLog.Debug("Updating user: ", id)
	updated, err := updateuu.UserDataInterface.UpdateById(id, user)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

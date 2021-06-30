package listuser

import (
	"github.com/muka-id/service-user/app/logger"
	"github.com/muka-id/service-user/applicationservice/dataservice"
	"github.com/muka-id/service-user/domain/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListUserUseCase struct {
	UserDataInterface dataservice.UserDataInterface
}

func (luc *ListUserUseCase) GetUserById(id string) (*model.User, error) {
	bsonId, err := primitive.ObjectIDFromHex(id)
	logger.SugarLog.Debugf("Find: %s", id)
	if err != nil {
		return nil, errors.Wrap(err, "Error converting string to bson "+id)
	}

	find, err := luc.UserDataInterface.Find(bsonId)
	if err != nil {
		return nil, err
	}
	return find, nil
}

func (luc *ListUserUseCase) GetUsers(filters interface{}) ([]*model.User, error) {
	return nil, nil
}

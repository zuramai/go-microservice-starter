package usecase

import (
	"github.com/muka-id/service-user/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetUserUseCaseInterface interface {
	GetUserById(id string) (*model.User, error)
	GetUsers(filters interface{}) ([]*model.User, error)
}

type RegistrationUseCaseInterface interface {
	RegisterUser(*model.User) (*model.User, error)
}

type UpdateUserUseCaseInterface interface {
	UpdateUserById(id primitive.ObjectID, user *model.User) (*model.User, error)
}

type DeleteUserUseCaseInterface interface {
	SoftDelete(id string) (rowsAffected int, err error)
	Delete(id string) (rowsAffected int, err error)
}

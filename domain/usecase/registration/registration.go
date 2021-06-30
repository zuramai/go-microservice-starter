package registration

import (
	"github.com/muka-id/service-user/applicationservice/dataservice"
	"github.com/muka-id/service-user/domain/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegistrationUseCase struct {
	UserDataInterface dataservice.UserDataInterface
}

// ErrUserExist is returned by isDuplicate methods when the user is exist in RegisterUser
var ErrUserExist = errors.New("cannot register, user exist")

func (ruc *RegistrationUseCase) RegisterUser(user *model.User) (*model.User, error) {
	err := user.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid fields")
	}
	isDuplicate, err := ruc.isDuplicate(user)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	if isDuplicate {
		return nil, ErrUserExist
	}

	resultUser, err := ruc.UserDataInterface.Insert(user)

	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return resultUser, nil
}

func (ruc *RegistrationUseCase) isDuplicate(user *model.User) (bool, error) {
	user, err := ruc.UserDataInterface.Find(user.Id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// Error, but the error is because there's no data.
			return false, nil
		} else {
			return false, errors.Wrap(err, "")
		}
	}
	if user != nil {
		return true, nil
	}
	return false, nil
}

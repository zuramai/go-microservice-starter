package userclient

import (
	"github.com/pkg/errors"

	"github.com/golang/protobuf/ptypes"
	"github.com/muka-id/service-user/app/logger"
	"github.com/muka-id/service-user/applicationservice/client/userclient/generatedclient"
	"github.com/muka-id/service-user/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConvertOptions struct {
	// if WithTimestamp false, when converting Grpc to User it doesnt create new timestamp
	WithTimestamp bool
}

func UserToGrpc(user *model.User) (*generatedclient.User, error) {
	if user == nil {
		return nil, nil
	}

	resultUser := generatedclient.User{}

	// Convert ObjectID to string
	resultUser.Id = user.Id.Hex()

	resultUser.Name = user.Name

	createdAt, err := ptypes.TimestampProto(user.CreatedAt)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	updatedAt, err := ptypes.TimestampProto(user.UpdatedAt)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	resultUser.CreatedAt = createdAt
	resultUser.UpdatedAt = updatedAt

	return &resultUser, nil
}
func GrpcToUser(user *generatedclient.User) (*model.User, error) {
	if user == nil {
		return nil, errors.New("Cannot convert nil values user")
	}

	resultUser := model.User{}

	// Convert ID String to OBjectId
	if len(user.Id) != 0 {
		userid, err := primitive.ObjectIDFromHex(user.Id)
		if err != nil {
			return nil, err
		}
		resultUser.Id = userid
	}
	logger.SugarLog.Debugf("Convert %s to %s", user.GetId(), resultUser.Id)
	resultUser.Name = user.GetName()

	// parse CreatedAt if exist
	if user.CreatedAt != nil {
		createdAt, err := ptypes.Timestamp(user.CreatedAt)
		if err != nil {
			return nil, errors.Wrap(err, "Error creating createdAt timestamp")
		}
		resultUser.CreatedAt = createdAt
	}
	// parse UpdatedAt if exist
	if user.UpdatedAt != nil {
		updatedAt, err := ptypes.Timestamp(user.UpdatedAt)
		if err != nil {
			return nil, errors.Wrap(err, "Error creating updatedAt timestamp")
		}
		resultUser.UpdatedAt = updatedAt
	}

	resultUser.Role = user.Role

	return &resultUser, nil

}

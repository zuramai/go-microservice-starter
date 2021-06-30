package deleteuser

import (
	"github.com/muka-id/service-user/applicationservice/dataservice"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteUserUseCase struct {
	UserDataInterface dataservice.UserDataInterface
}

// Delete method is permanently delete data
func (duc *DeleteUserUseCase) Delete(id string) (int32, error) {
	// Convert id string to ObjectId
	objectid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return 0, status.Error(codes.InvalidArgument, "Invalid id")
	}

	rowsAffected, err := duc.UserDataInterface.DeleteById(objectid)

	if err != nil {
		return 0, status.Error(codes.Internal, "Failed to execute delete")
	}

	return rowsAffected, nil
}

// SoftDelete is a way to delete without actually deleting it
// It just update the DeletedAt value
func (duc *DeleteUserUseCase) SoftDelete(id string) (int32, error) {
	return 0, nil
}

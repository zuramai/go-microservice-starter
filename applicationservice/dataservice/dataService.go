package dataservice

import (
	"github.com/muka-id/service-user/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDataInterface interface {
	// Remove: deletes a user by user name from database
	DeleteById(id primitive.ObjectID) (int32, error)

	// Find: retrieves a user from database by user's id
	Find(id primitive.ObjectID) (*model.User, error)

	// FindByName: retrieves a user from database by user's name
	// FindByName(name string) (user *model.User, err error)

	// FindAll: retrieves all users from database as array of model.User
	// FindAll() ([]model.User, error)

	// Update: make changes to user data in database
	UpdateById(id primitive.ObjectID, user *model.User) (*model.User, error)

	// Insert: add user to database and return inserted user
	Insert(user *model.User) (resultUser *model.User, err error)
}

// CacheDataInterface represents interface for cache service, which is a micro-service
type CacheDataInterface interface {
	// Get handles call to Get function on Cache service
	Get(key string) ([]byte, error)
	// Store handles call to Store function on Cache service
	Store(key string, value []byte) error
}

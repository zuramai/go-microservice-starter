package mongodb

import (
	"context"
	"errors"
	"time"

	"github.com/muka-id/service-user/app/logger"
	"github.com/muka-id/service-user/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserDataMongo struct {
	Connection *mongo.Client
	DB         *mongo.Database
}

func (usermongo *UserDataMongo) Find(id primitive.ObjectID) (*model.User, error) {
	logger.SugarLog.Debug("User find ", id)
	ctx := context.Background()
	collection := usermongo.DB.Collection("users")

	var user model.User

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.SugarLog.Debug("User not found")
		} else {
			logger.SugarLog.Debug("Find() Error: ", err)
		}
		return nil, err
	}

	logger.SugarLog.Debug("User found - ", user)

	return &user, nil
}

// func (usermongo *UserDataMongo) FindByName(name string) (*model.User, error) {

// }

func (usermongo *UserDataMongo) FindAll() ([]*model.User, error) {
	return nil, nil
}

func (usermongo *UserDataMongo) UpdateById(id primitive.ObjectID, user *model.User) (*model.User, error) {
	ctx := context.Background()

	collection := usermongo.DB.Collection("users")
	user.UpdatedAt = time.Now()
	userbson, err := bson.Marshal(user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error marhsaling bson")
	}

	result := collection.FindOneAndUpdate(ctx, bson.M{"_id": id}, userbson)

	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return nil, result.Err()
	}

	// Decode the SingleResult
	var updatedUser model.User
	result.Decode(&updatedUser)

	logger.SugarLog.Debug("Result update: ", updatedUser)

	return &updatedUser, nil
}

func (usermongo *UserDataMongo) Insert(user *model.User) (*model.User, error) {
	ctx := context.Background()
	collection := usermongo.DB.Collection("users")

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	userbson, err := bson.Marshal(user)

	if err != nil {
		return nil, err
	}

	inserted, err := collection.InsertOne(ctx, userbson)
	if err != nil {
		logger.SugarLog.Error("Insert error", err)
		return nil, err
	}

	logger.SugarLog.Debug("Insert Success - ", inserted.InsertedID)

	user.Id = inserted.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (usermongo *UserDataMongo) DeleteById(id primitive.ObjectID) (int32, error) {
	ctx := context.Background()

	collection := usermongo.DB.Collection("users")

	result := collection.FindOneAndDelete(ctx, bson.M{"_id": id})

	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return 0, result.Err()
	}

	return 1, nil

}

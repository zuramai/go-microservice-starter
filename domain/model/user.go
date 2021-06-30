package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Role      string             `bson:"role"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

// Validate validates a newly created user, which has not persisted to database yet, so Id is empty
func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required))
}

//ValidatePersisted validate a user that has been persisted to database, basically Id is not empty
func (u User) ValidatePersisted() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Id, validation.Required),
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.CreatedAt, validation.Required))
}

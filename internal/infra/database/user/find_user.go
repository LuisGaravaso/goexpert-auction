package user

import (
	"context"
	"errors"

	"github.com/LuisGaravaso/goexpert-auction/configs/logger"
	"github.com/LuisGaravaso/goexpert-auction/internal/entity/user"
	"github.com/LuisGaravaso/goexpert-auction/internal/internal_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongo struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: database.Collection("user"),
	}
}

func (r *UserRepository) FindUserById(ctx context.Context, id string) (*user.User, *internal_errors.InternalError) {
	filter := bson.M{"_id": id}
	var userMongo UserMongo
	if err := r.Collection.FindOne(ctx, filter).Decode(&userMongo); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error("user not found with id: "+id, err)
			return nil, internal_errors.NewNotFoundError("user not found with id: " + id)
		}

		logger.Error("error finding user by id: "+id, err)
		return nil, internal_errors.NewInternalServerError("error finding user by id: " + id)
	}

	return &user.User{
		ID:   userMongo.ID,
		Name: userMongo.Name,
	}, nil
}

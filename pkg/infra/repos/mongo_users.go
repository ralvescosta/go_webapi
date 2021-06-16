package repos

import (
	"context"
	"fmt"

	"webapi/pkg/app/dtos"
	"webapi/pkg/app/entities"
	"webapi/pkg/app/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepository struct {
	db *mongo.Client
}

func (r mongoUserRepository) Create(ctx context.Context, user *dtos.UserDto) (*entities.User, error) {
	collection := r.db.Database("GO").Collection("users")
	res, err := collection.InsertOne(ctx,
		bson.D{
			{"first_name", user.FirstName},
			{"last_name", user.LastName},
			{"email", user.Email},
			{"password", user.Password},
		})
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return &entities.User{}, nil
}

func (r mongoUserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	return &entities.User{}, nil
}

func NewUserMongoRepository(db *mongo.Client) interfaces.IUserRepository {
	return &mongoUserRepository{db: db}
}

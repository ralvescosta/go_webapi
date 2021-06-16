package repos

import (
	"context"
	"fmt"
	"time"

	"webapi/pkg/app/dtos"
	"webapi/pkg/app/entities"
	"webapi/pkg/app/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepository struct {
	db *mongo.Database
}

func (r mongoUserRepository) Create(ctx context.Context, user *dtos.UserDto) (*entities.User, error) {
	collection := r.db.Collection("users")
	res, err := collection.InsertOne(ctx,
		bson.D{
			{"first_name", user.FirstName},
			{"last_name", user.LastName},
			{"email", user.Email},
			{"password", user.Password},
			{"created_at", time.Now()},
			{"updated_at", time.Now()},
			{"deleted_at", nil},
		})
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return &entities.User{}, nil
}

func (r mongoUserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	collection := r.db.Collection("users")

	user := &entities.User{}
	err := collection.FindOne(ctx, bson.D{{"email", email}}).Decode(user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserMongoRepository(db *mongo.Database) interfaces.IUserRepository {
	return &mongoUserRepository{db: db}
}

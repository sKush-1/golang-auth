package user

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo struct {
	col *mongo.Collection
}

func NewRepo(db *mongo.Database) *Repo {
	return &Repo{col: db.Collection("users")}
}

func (r *Repo) EnsureEmailUniqueIndex(ctx context.Context) error {
	_, err := r.col.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "email", Value: 1}},
		Options: options.Index().
			SetName("email_unique").
			SetUnique(true).
			SetSparse(true),
	})
	if err != nil {
		return fmt.Errorf("failed to ensure unique user email index: %w", err)
	}
	return nil
}

func (r *Repo) Create(ctx context.Context, u User) (User, error) {
	res, err := r.col.InsertOne(ctx, u)
	if err != nil {
		return User{}, err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return User{}, fmt.Errorf("inserted user ID has type %T, expected primitive.ObjectID", res.InsertedID)
	}

	u.ID = id
	return u, nil
}

func (r *Repo) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User

	err := r.col.FindOne(ctx, bson.M{
		"email": strings.TrimSpace(email),
	}).Decode(&user)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

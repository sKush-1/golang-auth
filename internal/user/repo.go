package user

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	col *mongo.Collection
}

func NewRepo(db *mongo.Database) *Repo {
	return &Repo{col: db.Collection("users")}
}

func (r *Repo) Create(ctx context.Context, u User) (User, error) {
	res, err := r.col.InsertOne(ctx, u)
	if err != nil {
		return User{}, err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return User{}, fmt.Errorf("id is not objectId", err)
	}

	u.ID = id
	return u, nil
}

func (r *Repo) FindByEmail(ctx context.Context, email string) (User, error) {
	email = strings.TrimSpace(email)
	var user User
	err := r.col.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

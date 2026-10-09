package user

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct{
	col *mongo.Collection
}

func NewRepo(db *mongo.Database) *Repo {
	return &Repo(col: db.Collection("users"))
}

func (r *Repo) Create(ctx context.Context, u User) (User, error){
	res, err := r.col.InsertOne(ctx, u)
	if err != nil {
		return User{}, err
	}

	u.ID = res.InsertedID.(primitive.ObjectID)
	return u, nil
}
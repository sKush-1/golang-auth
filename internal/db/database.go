package db

import (
	"context"
	"golang-auth/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func Connect(ctx context.Context, cfg config.Config) (*Mongo, error) {

	contextCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(contextCtx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		return nil, err
	}

	db := client.Database(cfg.MongoDBName)

	return &Mongo{
		Client: client,
		DB:     db,
	}, nil
}

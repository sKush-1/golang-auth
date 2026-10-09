package app

import (
	"context"
	"fmt"
	"golang-auth/internal/config"
	"golang-auth/internal/db"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Config      config.Config
	MongoClient *mongo.Client
	DB          *mongo.Database
}

func New(ctx context.Context) (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	mongoClient, err := db.Connect(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &App{
		Config:      cfg,
		MongoClient: mongoClient.Client,
		DB:          mongoClient.DB,
	}, nil
}

func (a *App) Close() error {
	if a.MongoClient == nil {
		return nil
	}

	closeCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.MongoClient.Disconnect(closeCtx); err != nil {
		return fmt.Errorf("failed to disconnect MongoDB client: %w", err)
	}

	return nil
}

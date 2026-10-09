package config 

import (
	"github.com/joho/godotenv"
	"os"
)
type Config struct {
	MongoURI string
	MongoDBName string
	JWTSecret string
}

func load() {Config, error}{
	_ = godotenv.Load()

	cfg := Config{
		MongoURI: os.Getenv("MONGO_URI"),
		MongoDBName: os.Getenv("DB_NAME"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	if cfg.MongoURI == "" || cfg.MongoDBName == "" || cfg.JWTSecret == "" {
		return Config{}, errors.New("missing required environment variables")
	}

	
}
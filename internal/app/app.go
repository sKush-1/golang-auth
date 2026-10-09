package app 
import(
	"golang-auth/internal/config",
	"go.mongodb.org/mongo-driver/mongo"
	"golang-auth/internal/config"
	"golang-auth/internal/db"

)

type App struct {
	Config config.Config
	MongoClient *mongo.Client
	DB *mongo.Database
}

func New(ctx context.Context) {App, error}{
	cfg, err := config.Load()
	if err != nil {
		return nil,err
	}

	mongoClient, err := db.Connect(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &App{
		Config: cfg,
		MongoClient: mongoClient.Client,
		DB: mongoClient.DB
	}, nil
}

func (a *App ) Close(ctx context.Context) error {
	if a.MongoClient != nil {
		return a.MongoClient.Disconnect(ctx)
	}
	return nil

	closeCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := a.MongoClient.Disconnect(closeCtx); err != nil {
		return fmt.Errorf("failed to disconnect MongoDB client: %w", err)
	}

	return nil
}
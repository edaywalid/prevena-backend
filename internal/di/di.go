package di

import (
	"context"
	"time"

	"github.com/edaywalid/pinktober-hackathon-backend/internal/config"
	"github.com/edaywalid/pinktober-hackathon-backend/internal/handlers"
	"github.com/edaywalid/pinktober-hackathon-backend/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Container struct {
	Services     *Services
	Repositories *Repositories
	Handlers     *Handlers
	Config       *config.Config
	Databases    *Databases
	Logger       *logger.MyLogger
}

type (
	Services     struct{}
	Repositories struct{}
	Handlers     struct {
		PingHandler *handlers.PingHandler
		PingHandler    *handlers.PingHandler
	}
	Databases struct {
		DB *mongo.Database
	}
)

func NewContainer(log *logger.MyLogger) (*Container, error) {
	var container Container
	container.Logger = log

	container.Logger.LogInfo().Msg("Loading config")
	config, err := config.LoadConfig()
	if err != nil {
		container.Logger.LogError().Msgf("Error loading config: %v", err)
		return nil, err
	}
	container.Logger.LogInfo().Msg("Config loaded successfully")
	container.Config = config

	container.Logger.LogInfo().Msg("Initializing databases")
	if err := container.initDatabases(); err != nil {
		log.LogError().Msgf("Error initializing databases: %v", err)
		return nil, err
	}
	log.LogInfo().Msg("Databases initialized successfully")

	log.LogInfo().Msg("Initializing services, repositories, and handlers")
	container.InitServices()
	log.LogInfo().Msg("Services initialized successfully")
	container.InitRepositories()
	log.LogInfo().Msg("Repositories initialized successfully")
	container.InitHandlers()
	log.LogInfo().Msg("Handlers initialized successfully")

	return &container, nil
}

func (c *Container) initDatabases() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.Config.MONGO_URI))

	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}
	c.Databases = &Databases{}
	c.Databases.DB = client.Database(c.Config.DB_NAME)
	return nil
}

func (c *Container) InitServices() {
	services := &Services{}
	c.Services = services

}

func (c *Container) InitRepositories() {
	repositories := &Repositories{}
	c.Repositories = repositories
}

func (c *Container) InitHandlers() {
	handlers := &Handlers{
		PingHandler: handlers.NewPingHandler(),
		PingHandler:    handlers.NewPingHandler(),
	}
	c.Handlers = handlers
}

func (c *Container) Close() error {
	c.Logger.LogInfo().Msg("Closing databases")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if c.Databases.DB != nil {
		return c.Databases.DB.Client().Disconnect(ctx)
	}

	return nil
}

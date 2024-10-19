package di

import (
	"context"
	"time"
	"github.com/edaywalid/pinktober-hackathon-backend/internal/config"
	"github.com/edaywalid/pinktober-hackathon-backend/internal/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Container struct {
	Services     *Services
	Repositories *Repositories
	Handlers     *Handlers
	Config       *config.Config
	Databases    *Databases
}

type (
	Services     struct{}
	Repositories struct{}
	Handlers     struct {
		PingHandler *handlers.PingHandler
	}
	Databases struct {
		DB *mongo.Database
	}
)

func NewContainer() (*Container, error) {
	var container Container
	config, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	container.Config = config
	container.Services = &Services{}
	container.Repositories = &Repositories{}
	container.Handlers = &Handlers{
		PingHandler: &handlers.PingHandler{},
	if err := container.initDatabases(); err != nil {
		return nil, err
	}

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

func (c *Container) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if c.Databases.DB != nil {
		return c.Databases.DB.Client().Disconnect(ctx)
	}

	return nil
}

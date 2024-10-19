package di

import (
	"context"
	"time"

	"github.com/edaywalid/pinktober-hackathon-backend/internal/cache"
	"github.com/edaywalid/pinktober-hackathon-backend/internal/config"
	"github.com/edaywalid/pinktober-hackathon-backend/internal/handlers"
	"github.com/edaywalid/pinktober-hackathon-backend/internal/repositories"
	"github.com/edaywalid/pinktober-hackathon-backend/internal/seed"
	"github.com/edaywalid/pinktober-hackathon-backend/internal/services"
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
	Cache        *Cache
}

type (
	Services struct {
		ProductService *services.ProductService
	}
	Repositories struct {
		productRepository    *repositories.ProductRepository
		ingredientRepository *repositories.IngredientRepository
	}
	Handlers struct {
		PingHandler    *handlers.PingHandler
		SwaggerHandler *handlers.SwaggerHandler
		ProductHandler *handlers.ProductHandler
		SeedHandler    *handlers.SeedHandler
	}
	Databases struct {
		DB *mongo.Database
	}
	Cache struct {
		redis *cache.Redis
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

	container.InitRepositories()
	log.LogInfo().Msg("Repositories initialized successfully")

	container.InitServices()
	log.LogInfo().Msg("Services initialized successfully")

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
	services := &Services{
		ProductService: services.NewProductService(
			c.Repositories.productRepository,
			c.Repositories.ingredientRepository,
		),
	}
	c.Services = services
}

func (c *Container) InitCache() {
	redis, err := cache.NewRedis("localhost:6379")
	if err != nil {
		c.Logger.LogError().Msgf("Error initializing redis: %v", err)
		return
	}
	c.Cache = &Cache{
		redis: redis,
	}
}

func (c *Container) InitRepositories() {
	p_documents := c.Databases.DB.Collection("products")
	i_documents := c.Databases.DB.Collection("ingredients")
	repositories := &Repositories{
		productRepository: repositories.NewProductRepository(
			p_documents,
			i_documents,
		),
		ingredientRepository: repositories.NewIngredientRepository(i_documents),
	}
	c.Repositories = repositories
}

func (c *Container) InitHandlers() {
	handlers := &Handlers{
		PingHandler: handlers.NewPingHandler(),
		SwaggerHandler: handlers.NewSwaggerHandler(
			c.Config,
			c.Logger,
		),
		ProductHandler: handlers.NewProductHandler(
			c.Services.ProductService,
		),
		SeedHandler: handlers.NewSeedHandler(
			seed.NewSeeder(
				c.Repositories.productRepository,
				c.Repositories.ingredientRepository,
			),
		),
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

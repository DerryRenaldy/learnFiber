package server

import (
	"database/sql"
	"fmt"
	"github.com/DerryRenaldy/learnFiber/pkg/database"
	"log"
	"time"

	handlerV1 "github.com/DerryRenaldy/learnFiber/api/v1/handlers"
	serviceV1 "github.com/DerryRenaldy/learnFiber/api/v1/services"
	handlerV2 "github.com/DerryRenaldy/learnFiber/api/v2/handlers"
	serviceV2 "github.com/DerryRenaldy/learnFiber/api/v2/services"
	redisClients "github.com/DerryRenaldy/learnFiber/pkg/redis"
	"github.com/DerryRenaldy/learnFiber/server/middleware"
	"github.com/DerryRenaldy/learnFiber/store/mysql/customer"
	customercache "github.com/DerryRenaldy/learnFiber/store/redis/customer"
	"github.com/DerryRenaldy/logger/logger"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	logger    logger.ILogger
	serviceV1 serviceV1.IService
	handlerV1 handlerV1.CustomersHandlerInterface
	serviceV2 serviceV2.IService
	handlerV2 handlerV2.CustomersHandlerInterface
}

var SVR *Server
var db *sql.DB
var redisClient *redis.Client

func (s *Server) Register() {
	// MYSQL
	dbconnection := database.NewDatabaseConnection(s.logger)
	db = dbconnection.DBConnect()
	if db == nil {
		s.logger.Fatal("Expecting db connection object but received nil")
	}

	// REDIS
	redisConnection := redisClients.NewRedisConnection(s.logger)
	redisClient = redisConnection.RedisConnect()
	if redisClient == nil {
		s.logger.Fatal("Expecting db connection object but received nil")
	}
	redisCacheObj := customercache.NewRedis(redisClient, "fiber-project", 60*time.Second, s.logger)

	customerStore := customer.NewCustomerStoreImpl(s.logger, db)

	s.serviceV1 = serviceV1.New(customerStore, s.logger, redisCacheObj)
	s.serviceV2 = serviceV2.New(customerStore, s.logger, redisCacheObj)

	s.handlerV1 = handlerV1.NewCustomerHttpHandler(s.logger, s.serviceV1)
	s.handlerV2 = handlerV2.NewCustomerHttpHandler(s.logger, s.serviceV2)
}

func New(logger logger.ILogger) *Server {
	if SVR != nil {
		return SVR
	}
	SVR = &Server{
		logger: logger,
	}

	SVR.Register()

	return SVR
}

func (s Server) Start() {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})
	v1 := app.Group("/api/v1")
	v1.Use(recover.New())
	v1.Use(middleware.ValidateHeaderMiddleware())

	v1.Get("/", s.handlerV1.GetCustomerHandler)

	v2 := app.Group("/api/v2")
	v2.Use(recover.New())
	v2.Use(middleware.ValidateHeaderMiddleware())
	v2.Use(idempotency.New(
		idempotency.Config{
			KeyHeader: "X-Idempotency-Key",
		}))
	v2.Post("/", s.handlerV2.CreateCustomerHandler)
	v2.Post("/test", func(ctx *fiber.Ctx) error {
		ctx.Set("X-My-Header", "Hello from middleware")
		return ctx.SendString(fmt.Sprintf("header is: %s", ctx.Get("X-Idempotency-Key")))
	})

	log.Print("Hello from Cloud Run! The container started successfully and is listening for HTTP requests on $PORT")

	app.Listen(":3000")
}

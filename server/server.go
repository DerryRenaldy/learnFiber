package server

import (
	"database/sql"
	handlerV1 "github.com/DerryRenaldy/learnFiber/api/v1/handlers"
	serviceV1 "github.com/DerryRenaldy/learnFiber/api/v1/services"
	handlerV2 "github.com/DerryRenaldy/learnFiber/api/v2/handlers"
	serviceV2 "github.com/DerryRenaldy/learnFiber/api/v2/services"
	"github.com/DerryRenaldy/learnFiber/pkg/database"
	"github.com/DerryRenaldy/learnFiber/server/middleware"
	"github.com/DerryRenaldy/learnFiber/store/mysql/customer"
	"github.com/DerryRenaldy/logger/logger"
	"github.com/gofiber/fiber/v2"
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

func (s *Server) Register() {
	//MYSQL
	dbconnection := database.NewDatabaseConnection(s.logger)
	db = dbconnection.DBConnect()
	if db == nil {
		s.logger.Fatal("Expecting db connection object but received nil")
	}

	customerstore := customer.NewCustomerStoreImpl(s.logger, db)

	s.serviceV1 = serviceV1.New(customerstore, s.logger)
	s.serviceV2 = serviceV2.New(customerstore, s.logger)

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
	app := fiber.New()
	v1 := app.Group("/api/v1")
	v1.Use(middleware.ValidateHeaderMiddleware())

	v1.Get("/", s.handlerV1.GetCustomerHandler)

	v2 := app.Group("/api/v2")
	v2.Use(middleware.ValidateHeaderMiddleware())
	v2.Post("/", s.handlerV2.CreateCustomerHandler)

	app.Listen(":3000")
}

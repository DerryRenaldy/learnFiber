package server

import (
	"database/sql"
	"github.com/DerryRenaldy/learnFiber/api/v1/handlers"
	"github.com/DerryRenaldy/learnFiber/api/v1/services"
	"github.com/DerryRenaldy/learnFiber/pkg/database"
	"github.com/DerryRenaldy/learnFiber/store/mysql/customer"
	"github.com/DerryRenaldy/logger/logger"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	logger  logger.ILogger
	service services.IService
	handler handlers.CustomersHandlerInterface
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

	s.service = services.New(customerstore, s.logger)

	s.handler = handlers.NewCustomerHttpHandler(s.logger, s.service)
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
	app := fiber.New(fiber.Config{ColorScheme: fiber.DefaultColors, DisableKeepalive: true, Prefork: true})
	app.Get("/api/v1/", s.handler.GetCustomerHandler)

	app.Listen(":3000")
}

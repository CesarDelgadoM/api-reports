package main

import (
	"github.com/CesarDelgadoM/api-reports/config"
	"github.com/CesarDelgadoM/api-reports/internal/api/branch"
	"github.com/CesarDelgadoM/api-reports/internal/api/restaurant"
	"github.com/CesarDelgadoM/api-reports/internal/api/user"
	"github.com/CesarDelgadoM/api-reports/internal/report"
	"github.com/CesarDelgadoM/api-reports/pkg/database"
	"github.com/CesarDelgadoM/api-reports/pkg/logger/zap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type server struct {
	config *config.Config
}

func newServer(cfg *config.Config) *server {
	return &server{
		config: cfg,
	}
}

func main() {
	// Config load
	loadcfg := config.LoadConfig("config-local.yml")
	cfg := config.ParseConfig(loadcfg)

	server := newServer(cfg)
	server.run()
}

func (s *server) run() {
	// Logger zap initialization
	zap.InitLogger(s.config)

	// Postgres connection
	postgresConn := database.ConnectPostgresDB(s.config.Postgres)

	// Mongo connection
	mongoConn := database.ConnectMongoDB(s.config.Mongo)
	defer mongoConn.Disconnect()

	// Fiber app
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: s.config.Cors.AllowCredentials,
		AllowOrigins:     s.config.Cors.AllowOrigins,
		AllowHeaders:     s.config.Cors.AllowHeaders,
	}))

	// User
	userRepository := user.NewRepository(postgresConn)
	userService := user.NewUserService(userRepository)
	user.NewHandler(app, userService)

	// Restaurant
	restaurantRepository := restaurant.NewRepository(mongoConn)
	restaurantService := restaurant.NewService(restaurantRepository)
	restaurant.NewHandler(app, restaurantService)

	// Branch
	branchRepository := branch.NewRepository(mongoConn)
	branchService := branch.NewService(branchRepository)
	branch.NewHandler(app, branchService)

	// Reports
	report.NewHandler(app, s.config.Services)

	app.Listen(s.config.Server.Port)
}

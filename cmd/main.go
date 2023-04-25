package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/work/SRM/api/V1/controllers"
	_ "github.com/work/SRM/api/docs"
	"github.com/work/SRM/configs"
	"github.com/work/SRM/pkg/middleware"
	"github.com/work/SRM/pkg/utils"
	"github.com/work/SRM/routes"
	"github.com/work/SRM/service"

	//register swagger
	"log"
)

var (
	fiberConfig = configs.FiberConfig()
	appConfig   = configs.Config()
)

// @title API
// @version 1
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name Soburjon
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New(fiberConfig)

	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	middleware.FiberMiddleware(app)

	jwtRoleAuthorizer, err := middleware.NewJWTRoleAuthorizer(appConfig)
	if err != nil {
		log.Fatal("Could not initialize JWT Role Authorizer")
	}

	psqlString, _ := utils.ConnectionURLBuilder("postgres")
	connDb, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		//mylogger.Error("postgres connect error",
		//	logger.Error(err))
		return
	}

	utils.MigrationsUp()

	adminService := service.NewAdminService(connDb)
	programmerService := service.NewProgrammerService(connDb)
	registerService := service.NewRegisterService(connDb)

	api := controllers.NewApi(adminService, programmerService, registerService)

	app.Use(middleware.NewAuthorizer(jwtRoleAuthorizer))
	routes.SwaggerRoute(app)
	routes.AdminRoutes(app, api)
	routes.RegisterRoutes(app, api)
	routes.ProgrammerRoutes(app, api)

	app.Listen(":8000")
}

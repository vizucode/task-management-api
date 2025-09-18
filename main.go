// @title Task Management API
// @version 1.0
// @description API untuk mengelola tasks dengan authentication JWT
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

package main

import (
	"task-management-api/apps/middlewares/security"
	"task-management-api/apps/repositories/psql"
	routerRest "task-management-api/apps/router/rest"
	"task-management-api/apps/service/task"
	"task-management-api/apps/service/user"
	errorhandler "task-management-api/helpers/error_handler"

	_ "task-management-api/docs" // Import docs

	"github.com/go-playground/validator/v10"
	"github.com/vizucode/gokit/adapter/dbc"
	"github.com/vizucode/gokit/config"
	"github.com/vizucode/gokit/factory/server"
	"github.com/vizucode/gokit/factory/server/rest"
	"github.com/vizucode/gokit/utils/constant"
	"github.com/vizucode/gokit/utils/env"
)

func main() {

	/*
		Library
	*/
	serviceName := env.GetString("SERVICE_NAME")
	config.Load(serviceName, ".")
	validator10 := validator.New()

	dbConnection := dbc.NewGormConnection(
		dbc.SetGormURIConnection(env.GetString("DB_CONNECTION")),
		dbc.SetGormDriver(constant.Postgres),
		dbc.SetGormMaxIdleConnection(2),
		dbc.SetGormMaxPoolConnection(50),
		dbc.SetGormMinPoolConnection(10),
		dbc.SetGormSkipTransaction(true),
		dbc.SetGormServiceName(serviceName),
	)

	/*
		Repositories
	*/
	postgreDB := psql.NewPsql(dbConnection.DB)

	/*
		Service Mapping
	*/
	taskService := task.NewTask(
		postgreDB,
		validator10,
	)

	userService := user.NewUser(
		postgreDB,
		validator10,
	)

	restRouter := routerRest.NewRest(
		security.NewSecurity(),
		taskService,
		userService,
	)

	app := server.NewService(
		server.SetServiceName(serviceName),
		server.SetRestHandler(restRouter),
		server.SetRestHandlerOptions(
			rest.SetHTTPHost(env.GetString("APP_HOST")),
			rest.SetHTTPPort(env.GetInteger("APP_PORT")),
			rest.SetErrorHandler(errorhandler.FiberErrHandler),
		),
	)

	appServer := server.New(app)
	appServer.Run()
}

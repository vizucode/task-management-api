package main

import (
	"github.com/origamilabsid/backend-boilerplate/apps/middlewares/security"
	"github.com/origamilabsid/backend-boilerplate/apps/repositories/psql"
	routerRest "github.com/origamilabsid/backend-boilerplate/apps/router/rest"
	"github.com/origamilabsid/backend-boilerplate/apps/service/todo"
	errorhandler "github.com/origamilabsid/backend-boilerplate/helpers/error_handler"

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

	// apiClient := request.NewRequest(&http.Client{
	// 	Timeout: 5 * time.Second,
	// })

	/*
		Service Mapping
	*/
	todoService := todo.NewTodoService(
		postgreDB,
		validator10,
	)

	restRouter := routerRest.NewRest(
		security.NewSecurity(),
		todoService,
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

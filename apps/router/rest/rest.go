package rest

import (
	"strconv"

	"task-management-api/apps/domain"
	"task-management-api/apps/middlewares"
	"task-management-api/apps/service"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

type rest struct {
	mw   middlewares.IMiddleware
	task service.ITaskService
	user service.IUserService
}

func NewRest(
	mw middlewares.IMiddleware,
	task service.ITaskService,
	user service.IUserService,
) *rest {
	return &rest{
		mw:   mw,
		task: task,
		user: user,
	}
}

func (r *rest) Router(app fiber.Router) {
	// Swagger documentation
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// User routes
	app.Post("/signup", r.SignUp)
	app.Post("/signin", r.SignIn)

	// Task routes (protected)
	app.Get("/tasks", r.mw.AuthMiddleware, r.GetTaskList)
	app.Get("/tasks/:id", r.mw.AuthMiddleware, r.GetTaskDetail)
	app.Post("/tasks", r.mw.AuthMiddleware, r.CreateTask)
	app.Put("/tasks", r.mw.AuthMiddleware, r.UpdateTask)
	app.Delete("/tasks/:id", r.mw.AuthMiddleware, r.DeleteTask)
}

func (rest *rest) ResponseJson(
	ctx *fiber.Ctx,
	StatusCode int,
	data interface{},
	metadata domain.Metadata,
	message string,
) error {
	return ctx.Status(StatusCode).JSON(&domain.ResponseJson{
		StatusCode: strconv.Itoa(StatusCode),
		Data:       data,
		Metadata:   metadata,
		Message:    message,
	})
}

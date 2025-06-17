package rest

import (
	"strconv"

	"github.com/origamilabsid/backend-boilerplate/apps/domain"
	"github.com/origamilabsid/backend-boilerplate/apps/middlewares"
	"github.com/origamilabsid/backend-boilerplate/apps/service"

	"github.com/gofiber/fiber/v2"
)

type rest struct {
	mw middlewares.IMiddleware

	todo service.ITodoService
}

func NewRest(
	mw middlewares.IMiddleware,
	todo service.ITodoService,
) *rest {
	return &rest{
		mw:   mw,
		todo: todo,
	}
}

func (r *rest) Router(app fiber.Router) {
	v1 := app.Group("/v1")

	v1.Post("/todo", r.mw.LangTranslate, r.mw.AuthMiddleware, r.getTodo)
}

func (rest *rest) ResponseJson(
	ctx *fiber.Ctx,
	StatusCode int,
	data interface{},
	message string,
) error {
	return ctx.Status(StatusCode).JSON(&domain.ResponseJson{
		StatusCode: strconv.Itoa(StatusCode),
		Data:       data,
		Message:    message,
	})
}

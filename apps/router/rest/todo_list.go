package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vizucode/gokit/logger"
)

func (r *rest) getTodo(c *fiber.Ctx) error {

	// get todo
	todo, err := r.todo.GetTodo(c.Context())
	if err != nil {
		logger.Log.Debug(c.Context(), "error get todo")
	}

	return r.ResponseJson(c, http.StatusOK, todo, "Successfully get todo")
}

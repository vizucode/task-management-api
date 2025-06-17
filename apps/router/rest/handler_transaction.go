package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/origamilabsid/backend-boilerplate/apps/middlewares/security"
	"github.com/vizucode/gokit/logger"
)

func (r *rest) getTodo(c *fiber.Ctx) error {
	ctx, _, ok := security.ExtractUserContextFiber(c)
	if !ok {
		logger.Log.Debug(ctx, "error extract user context")
	}

	// get todo
	todo, err := r.todo.GetTodo(ctx)
	if err != nil {
		logger.Log.Debug(ctx, "error get todo")
	}

	// translate
	// targetLang := c.Locals("lang").(string)
	// messsage, _ := translator.Translate("Successfully cancel transaction", "en", targetLang)

	return r.ResponseJson(c, http.StatusOK, todo, "Successfully get todo")
}

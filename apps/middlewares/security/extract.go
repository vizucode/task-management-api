package security

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/origamilabsid/backend-boilerplate/apps/domain"
	contextkeys "github.com/origamilabsid/backend-boilerplate/helpers/constants/context_keys"
)

func ExtractUserContextFiber(c *fiber.Ctx) (responseCtx context.Context, resp domain.UserContext, ok bool) {
	resultUserContex, ok := c.Locals(contextkeys.UserContext).(domain.UserContext)

	// set context with user info
	ctx := context.WithValue(c.UserContext(), contextkeys.UserContext, resultUserContex)
	return ctx, resultUserContex, ok
}

func ExtractUserContext(ctx context.Context) (resp domain.UserContext, ok bool) {
	resp, ok = ctx.Value(contextkeys.UserContext).(domain.UserContext)
	return resp, ok
}

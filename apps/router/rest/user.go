package rest

import (
	"net/http"

	"task-management-api/apps/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/vizucode/gokit/logger"
)

// SignUp godoc
// @Summary Register a new user
// @Description Register a new user with username, email, and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body domain.User true "User registration data"
// @Success 201 {object} domain.ResponseJson "User registered successfully"
// @Failure 400 {object} domain.ResponseJson "Bad request"
// @Router /signup [post]
func (r *rest) SignUp(c *fiber.Ctx) error {
	var payload domain.User

	// Parse request body
	if err := c.BodyParser(&payload); err != nil {
		logger.Log.Error(c.Context(), err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, domain.Metadata{}, "Invalid request body")
	}

	// Call user service
	err := r.user.SignUp(c.Context(), payload)
	if err != nil {
		logger.Log.Error(c.Context(), err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, domain.Metadata{}, err.Error())
	}

	return r.ResponseJson(c, http.StatusCreated, nil, domain.Metadata{}, "User registered successfully")
}

// SignIn godoc
// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body domain.User true "User login credentials"
// @Success 200 {object} domain.ResponseJson{data=domain.SignInResponse} "Login successful"
// @Failure 400 {object} domain.ResponseJson "Bad request"
// @Failure 401 {object} domain.ResponseJson "Unauthorized"
// @Router /signin [post]
func (r *rest) SignIn(c *fiber.Ctx) error {
	var payload domain.User

	// Parse request body
	if err := c.BodyParser(&payload); err != nil {
		logger.Log.Error(c.Context(), err)
		return r.ResponseJson(c, http.StatusBadRequest, nil, domain.Metadata{}, "Invalid request body")
	}

	// Call user service
	resp, err := r.user.SignIn(c.Context(), payload)
	if err != nil {
		logger.Log.Error(c.Context(), err)
		return r.ResponseJson(c, http.StatusUnauthorized, nil, domain.Metadata{}, err.Error())
	}

	return r.ResponseJson(c, http.StatusOK, resp, domain.Metadata{}, "Login successful")
}

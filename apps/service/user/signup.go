package user

import (
	"context"
	"task-management-api/apps/domain"
	"task-management-api/apps/models"

	"github.com/google/uuid"
	"github.com/vizucode/gokit/logger"
	"golang.org/x/crypto/bcrypt"
)

func (uc *user) SignUp(ctx context.Context, payload domain.User) (err error) {

	err = uc.validator.VarCtx(ctx, payload.Username, "required")
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	err = uc.validator.StructCtx(ctx, payload)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	err = uc.db.CreateUser(ctx, models.User{
		ID:       uuid.New(),
		Username: payload.Username,
		Email:    payload.Email,
		Password: string(passwordHashed),
	})
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return nil
}

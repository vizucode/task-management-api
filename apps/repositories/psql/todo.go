package psql

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/origamilabsid/backend-boilerplate/apps/models"
	"github.com/origamilabsid/backend-boilerplate/helpers/constants/rpcstd"
	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
	"gorm.io/gorm"
)

func (psql *psql) GetTodo(ctx context.Context, id uuid.UUID) (resp models.Todo, err error) {
	err = psql.db.WithContext(ctx).Model(&models.Todo{}).Where("id = ?", id).First(&resp).Error
	if err != nil {
		logger.Log.Error(ctx, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		return resp, errorkit.NewErrorStd(http.StatusInternalServerError, rpcstd.INTERNAL, "can't get todo")
	}

	return resp, nil
}

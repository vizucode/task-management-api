package psql

import (
	"context"
	"errors"
	"net/http"
	"task-management-api/apps/models"
	"task-management-api/helpers/constants/httpstd"

	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
	"gorm.io/gorm"
)

func (psql *psql) CreateUser(ctx context.Context, user models.User) (err error) {

	err = psql.db.WithContext(ctx).Model(&models.User{}).Create(&user).Error
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return nil
}

func (psql *psql) GetUserByEmail(ctx context.Context, email string) (resp models.User, err error) {
	err = psql.db.WithContext(ctx).Model(&models.User{}).Where("email = ?", email).First(&resp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		logger.Log.Error(ctx, err)
		return resp, errorkit.NewErrorStd(http.StatusInternalServerError, "", httpstd.InternalServerError)
	}

	return resp, nil
}

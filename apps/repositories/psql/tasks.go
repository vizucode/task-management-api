package psql

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"task-management-api/apps/models"

	"github.com/vizucode/gokit/logger"
	"gorm.io/gorm"
)

func (psql *psql) GetTaskList(ctx context.Context, filter models.Filter) (resp []models.Task, err error) {
	tx := psql.db.WithContext(ctx).Model(&models.Task{})

	if filter.Limit > 0 {
		tx = tx.Limit(filter.Limit)
	}

	if filter.Page > 0 {
		offset := (filter.Page - 1) * filter.Limit
		tx = tx.Offset(offset)
	}

	if !strings.EqualFold(filter.TaskStatus, "") {
		tx = tx.Where("status = ?", filter.TaskStatus)
	}

	if !strings.EqualFold(filter.UserId, "") {
		tx = tx.Where("user_id = ?", filter.UserId)
	}

	err = tx.Find(&resp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		logger.Log.Error(ctx, err)
		return resp, err
	}

	return resp, nil
}

func (psql *psql) CountTaskList(ctx context.Context, filter models.Filter) (resp int64, err error) {

	tx := psql.db.WithContext(ctx).Model(&models.Task{})

	if !strings.EqualFold(filter.TaskStatus, "") {
		tx = tx.Where("status = ?", filter.TaskStatus)
	}

	if !strings.EqualFold(filter.UserId, "") {
		tx = tx.Where("user_id = ?", filter.UserId)
	}

	err = tx.Count(&resp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		logger.Log.Error(ctx, err)
		return resp, err
	}

	return resp, nil
}

func (psql *psql) GetDetailTask(ctx context.Context, taskId string) (resp models.Task, err error) {
	fmt.Println("taskid: ", taskId)
	err = psql.db.WithContext(ctx).Model(&models.Task{}).Where("id = ?", taskId).First(&resp).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Log.Error(ctx, err)
			return resp, err
		}
	}
	return resp, nil
}

func (psql *psql) CreateTask(ctx context.Context, payload models.Task) (err error) {
	err = psql.db.WithContext(ctx).Model(&models.Task{}).Create(&payload).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Log.Error(ctx, err)
			return err
		}
	}
	return nil
}

func (psql *psql) UpdateTask(ctx context.Context, payload models.Task) (err error) {

	fields := []string{}

	if !strings.EqualFold(payload.Title, "") {
		fields = append(fields, "title")
	}

	if !strings.EqualFold(payload.Description, "") {
		fields = append(fields, "description")
	}

	if !strings.EqualFold(payload.Status, "") {
		fields = append(fields, "status")
	}

	fmt.Println(payload.Status)
	fmt.Println(payload.Id)
	fmt.Println(fields)

	err = psql.db.WithContext(ctx).Model(&models.Task{}).Select(fields).Where("id = ?", payload.Id).Updates(&payload).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Log.Error(ctx, err)
			return err
		}
	}
	return nil
}

func (psql *psql) DeleteTask(ctx context.Context, taskId string) (err error) {

	err = psql.db.WithContext(ctx).Model(&models.Task{}).Where("id = ?", taskId).Delete(&models.Task{}).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Log.Error(ctx, err)
			return err
		}
	}

	return nil
}

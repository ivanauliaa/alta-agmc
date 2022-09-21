package repository

import (
	"context"
	"fmt"

	"day6-hexagonal/internal/model"

	"gorm.io/gorm"
)

type UserInterface interface {
	Create(ctx context.Context, data model.User) error
	FindAll(ctx context.Context) ([]model.User, error)
	FindById(ctx context.Context, id uint) (model.User, error)
	Update(ctx context.Context, id uint, data map[string]interface{}) error
	Delete(ctx context.Context, id uint) error
	FindByEmail(ctx context.Context, email string) (model.User, error)
	VerifyUserOwner(ctx context.Context, id uint, owner uint) error
}

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) UserInterface {
	return &user{
		db: db,
	}
}

func (r *user) Create(ctx context.Context, data model.User) error {
	return r.db.WithContext(ctx).Model(&model.User{}).Create(&data).Error
}

func (r *user) FindAll(ctx context.Context) ([]model.User, error) {
	var users []model.User

	if err := r.db.WithContext(ctx).Model(&model.User{}).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *user) FindById(ctx context.Context, id uint) (model.User, error) {
	var user model.User

	err := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).First(&user).Error

	return user, err
}

func (r *user) Update(ctx context.Context, id uint, data map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Updates(data).Error
}

func (r *user) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error
}

func (r *user) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User

	err := r.db.WithContext(ctx).Model(&model.User{}).Where("email = ?", email).First(&user).Error

	return user, err
}

func (r *user) VerifyUserOwner(ctx context.Context, id uint, owner uint) error {
	var user model.User

	if err := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return nil
	}

	if user.ID != owner {
		return fmt.Errorf("restricted resource")
	}

	return nil
}

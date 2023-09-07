package user

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"my-contacts/entity"
	"my-contacts/port/repository"
)

type repo struct {
	DB *gorm.DB
}

func NewGormUserRepository(gormdb *gorm.DB) repository.UserRepository {
	return &repo{
		DB: gormdb,
	}
}

func (r *repo) Insert(ctx context.Context, user *entity.User) error {
	err := r.DB.WithContext(ctx).Save(user).Error
	if err != nil {
		return fmt.Errorf("insert -> %w", err)
	}

	return nil
}

func (r *repo) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := r.DB.WithContext(ctx).Save(user).Error
	if err != nil {
		return nil, fmt.Errorf("save -> %w", err)
	}

	return user, nil
}

func (r *repo) ListAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	err := r.DB.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("find -> %w", err)
	}

	return users, nil
}

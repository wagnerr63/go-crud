package repository

import (
	"context"
	"go-crud/entity"
)

type UserRepository interface {
	Insert(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) (*entity.User, error)
	ListAll(ctx context.Context) ([]entity.User, error)
}

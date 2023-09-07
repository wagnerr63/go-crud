package user

import (
	"context"
	"fmt"
	"go-crud/entity"
	"go-crud/port/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository repository.UserRepository
}

func New(repo repository.UserRepository) *Service {
	return &Service{
		repository: repo,
	}
}

func (s *Service) Create(ctx context.Context, dto *CreateDTO) (*entity.User, error) {
	createdUser, err := s.create(ctx, dto)
	if err != nil {
		err = fmt.Errorf("create -> %w", err)
		fmt.Println("ERROR: ", err)
		return nil, err
	}

	return createdUser, nil
}

func (s *Service) create(ctx context.Context, dto *CreateDTO) (*entity.User, error) {
	err := dto.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate -> %w", err)
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("GenerateFromPassword -> %w", err)
	}

	newUser := &entity.User{
		ID:        uuid.NewString(),
		Name:      dto.Name,
		Email:     dto.Email,
		Password:  string(encryptedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.repository.Insert(ctx, newUser)
	if err != nil {
		return nil, fmt.Errorf("insert -> %w", err)
	}

	return newUser, nil
}

func (s *Service) ListAll(ctx context.Context) ([]entity.User, error) {
	users, err := s.listAll(ctx)
	if err != nil {
		err = fmt.Errorf("listAll -> %w", err)
		fmt.Println("ERROR: ", err)
		return nil, err
	}

	return users, nil
}

func (s *Service) listAll(ctx context.Context) ([]entity.User, error) {
	users, err := s.repository.ListAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("listAll -> %w", err)
	}

	return users, nil
}

package user

import (
	"context"
	"day6-hexagonal/internal/dto"
	"day6-hexagonal/internal/factory"
	"day6-hexagonal/internal/model"
	"day6-hexagonal/internal/repository"
)

type ServiceInterface interface {
	Create(ctx context.Context, payload *dto.CreateUserRequest) (string, error)
	FindAll(ctx context.Context) (*[]dto.UserResponse, error)
	FindById(ctx context.Context, payload *dto.ByIdRequest) (*dto.UserResponse, error)
	Update(ctx context.Context, id uint, owner uint, payload *dto.UpdateUserRequest) (string, error)
	Delete(ctx context.Context, owner uint, payload *dto.ByIdRequest) (*dto.UserResponse, error)
}

type service struct {
	UserRepository repository.UserInterface
}

func NewService(f *factory.Factory) ServiceInterface {
	return &service{
		UserRepository: f.UserRepository,
	}
}

func (s *service) Create(ctx context.Context, payload *dto.CreateUserRequest) (string, error) {
	user := model.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}

	if err := s.UserRepository.Create(ctx, user); err != nil {
		return "", err
	}

	return "success", nil
}

func (s *service) FindAll(ctx context.Context) (*[]dto.UserResponse, error) {
	users, err := s.UserRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]dto.UserResponse, len(users))
	for i := range result {
		result[i] = dto.UserResponse{
			ID:    users[i].ID,
			Name:  users[i].Name,
			Email: users[i].Email,
		}
	}

	return &result, err
}

func (s *service) FindById(ctx context.Context, payload *dto.ByIdRequest) (*dto.UserResponse, error) {
	user, err := s.UserRepository.FindById(ctx, payload.ID)
	if err != nil {
		return nil, err
	}

	result := new(dto.UserResponse)
	result.ID = user.ID
	result.Name = user.Name
	result.Email = user.Email

	return result, nil
}

func (s *service) Update(ctx context.Context, id uint, owner uint, payload *dto.UpdateUserRequest) (string, error) {
	if err := s.UserRepository.VerifyUserOwner(ctx, id, owner); err != nil {
		return "", err
	}

	data := make(map[string]interface{})

	if payload.Name != nil {
		data["name"] = payload.Name
	}

	if payload.Email != nil {
		data["email"] = payload.Email
	}

	if payload.Password != nil {
		data["password"] = payload.Password
	}

	if err := s.UserRepository.Update(ctx, id, data); err != nil {
		return "", err
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, owner uint, payload *dto.ByIdRequest) (*dto.UserResponse, error) {
	if err := s.UserRepository.VerifyUserOwner(ctx, payload.ID, owner); err != nil {
		return nil, err
	}

	user, err := s.UserRepository.FindById(ctx, payload.ID)
	if err != nil {
		return nil, err
	}

	if err := s.UserRepository.Delete(ctx, payload.ID); err != nil {
		return nil, err
	}

	result := new(dto.UserResponse)
	result.ID = user.ID
	result.Name = user.Name
	result.Email = user.Email

	return result, nil
}

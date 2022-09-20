package book

import (
	"context"
	"day6-hexagonal/internal/dto"
	"day6-hexagonal/internal/factory"
	"day6-hexagonal/internal/model"
	"time"

	"day6-hexagonal/internal/repository"
)

type ServiceInterface interface {
	Create(ctx context.Context, payload *dto.CreateBookRequest) (string, error)
	FindAll(ctx context.Context) (*[]model.Book, error)
	FindById(ctx context.Context, payload *dto.ByUnixIDRequest) (*model.Book, error)
	Update(ctx context.Context, id int64, payload *dto.UpdateBookRequest) (string, error)
	Delete(ctx context.Context, payload *dto.ByUnixIDRequest) (string, error)
}

type service struct {
	BookRepository repository.BookInterface
}

func NewService(f *factory.Factory) ServiceInterface {
	return &service{
		BookRepository: f.BookRepository,
	}
}

func (s *service) Create(ctx context.Context, payload *dto.CreateBookRequest) (string, error) {
	book := model.Book{
		BookId: time.Now().Unix(),
		Title:  payload.Title,
		Author: payload.Author,
		Year:   payload.Year,
	}

	if err := s.BookRepository.Create(ctx, book); err != nil {
		return "", err
	}

	return "success", nil
}

func (s *service) FindAll(ctx context.Context) (*[]model.Book, error) {
	books, err := s.BookRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return &books, nil
}

func (s *service) FindById(ctx context.Context, payload *dto.ByUnixIDRequest) (*model.Book, error) {
	book, err := s.BookRepository.FindById(ctx, payload.ID)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *service) Update(ctx context.Context, id int64, payload *dto.UpdateBookRequest) (string, error) {
	data := make(map[string]interface{})

	if payload.Title != nil {
		data["title"] = payload.Title
	}

	if payload.Author != nil {
		data["author"] = payload.Author
	}

	if payload.Year != nil {
		data["year"] = payload.Year
	}

	if err := s.BookRepository.Update(ctx, id, data); err != nil {
		return "", err
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, payload *dto.ByUnixIDRequest) (string, error) {
	if err := s.BookRepository.Delete(ctx, payload.ID); err != nil {
		return "", err
	}

	return "success", nil
}

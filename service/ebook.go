package service

import (
	"context"
	"dtam-fund-cms-backend/domain/entities"
	"dtam-fund-cms-backend/domain/ports"
)

type EBookService struct {
	ebookRepository ports.EBookRepository
}

func NewEBookService(
	ebookRepository ports.EBookRepository,
) ports.EBookService {
	return &EBookService{
		ebookRepository: ebookRepository,
	}
}

func (eb *EBookService) CreateEBook(ctx context.Context, ebook *entities.Ebook, thumbnail, file *entities.FileObject) error {
	return nil
}
func (eb *EBookService) GetEBook(ctx context.Context, id string) (*entities.Ebook, error) {
	return nil, nil
}
func (eb *EBookService) GetEBookList(ctx context.Context, page, limit string) ([]*entities.Ebook, error) {
	return nil, nil
}
func (eb *EBookService) EditEBook(ctx context.Context, id string, ebook *entities.Ebook, thumbnail, file *entities.FileObject) error {
	return nil
}
func (eb *EBookService) DeleteEBook(ctx context.Context, id string) error {
	return nil
}

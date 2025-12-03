package ports

import (
	"context"
	"dtam-fund-cms-backend/domain/entities"
)

type EBookRepository interface {
	SaveEBook(ctx context.Context, ebook *entities.Ebook) error
	RetriveEBook(ctx context.Context, id string) (*entities.Ebook, error)
	RetriveEBookList(ctx context.Context, page, limit int64) ([]*entities.Ebook, error)
	EditEBook(ctx context.Context, id string, ebook *entities.Ebook) error
	DeleteEBook(ctx context.Context, id string) error
}

type EBookService interface {
	CreateEBook(ctx context.Context, ebook *entities.Ebook, thumbnail, file *entities.FileObject) error
	GetEBook(ctx context.Context, id string) (*entities.Ebook, error)
	GetEBookList(ctx context.Context, page, limit string) ([]*entities.Ebook, error)
	EditEBook(ctx context.Context, id string, ebook *entities.Ebook, thumbnail, file *entities.FileObject) error
	DeleteEBook(ctx context.Context, id string) error
}

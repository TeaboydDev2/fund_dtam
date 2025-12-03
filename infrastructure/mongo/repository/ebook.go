package repository

import (
	"context"
	"dtam-fund-cms-backend/domain/entities"
	"dtam-fund-cms-backend/domain/ports"
	mongodb "dtam-fund-cms-backend/infrastructure/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

type EBookRepository struct {
	collection *mongo.Collection
}

func NewEBookRepository(db *mongodb.MongoClient) ports.EBookRepository {
	return &EBookRepository{
		collection: db.Collection("ebook"),
	}
}

func (eb *EBookRepository) SaveEBook(ctx context.Context, ebook *entities.Ebook) error {
	return nil
}
func (eb *EBookRepository) RetriveEBook(ctx context.Context, id string) (*entities.Ebook, error) {
	return nil, nil
}
func (eb *EBookRepository) RetriveEBookList(ctx context.Context, page, limit int64) ([]*entities.Ebook, error) {
	return nil, nil
}
func (eb *EBookRepository) EditEBook(ctx context.Context, id string, ebook *entities.Ebook) error {
	return nil
}
func (eb *EBookRepository) DeleteEBook(ctx context.Context, id string) error {
	return nil
}

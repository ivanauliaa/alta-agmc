package repository

import (
	"context"
	"day6-hexagonal/internal/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookInterface interface {
	Create(ctx context.Context, data model.Book) error
	FindAll(ctx context.Context) ([]model.Book, error)
	FindById(ctx context.Context, id int64) (model.Book, error)
	Update(ctx context.Context, id int64, data map[string]interface{}) error
	Delete(ctx context.Context, id int64) error
}

type book struct {
	db *mongo.Database
}

func NewBook(db *mongo.Database) BookInterface {
	return &book{
		db: db,
	}
}

func (r *book) Create(ctx context.Context, data model.Book) error {
	_, err := r.db.Collection("books").InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (r *book) FindAll(ctx context.Context) ([]model.Book, error) {
	csr, err := r.db.Collection("books").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer csr.Close(ctx)

	result := []model.Book{}
	for csr.Next(ctx) {
		row := model.Book{}
		err := csr.Decode(&row)
		if err != nil {
			return nil, err
		}

		result = append(result, row)
	}

	return result, nil
}

func (r *book) FindById(ctx context.Context, id int64) (model.Book, error) {
	var book model.Book
	var result map[string]interface{}

	err := r.db.Collection("books").FindOne(ctx, bson.M{"bookId": id}).Decode(&result)
	if err != nil {
		return book, err
	}

	book.BookId = result["bookId"].(int64)
	book.Title = result["title"].(string)
	book.Author = result["author"].(string)
	book.Year = int(result["year"].(int32))

	return book, nil
}

func (r *book) Update(ctx context.Context, id int64, data map[string]interface{}) error {
	selector := bson.M{
		"bookId": id,
	}

	changes := bson.M{
		"$set": data,
	}

	result, err := r.db.Collection("books").UpdateOne(ctx, selector, changes)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("book not found")
	}

	return nil
}

func (r *book) Delete(ctx context.Context, id int64) error {
	selector := bson.M{
		"bookId": id,
	}

	result, err := r.db.Collection("books").DeleteOne(ctx, selector)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("book not found")
	}

	return nil
}

package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/db/domain/model"
	"gorm.io/gorm"
	"os"
)

var LikeOperator = map[string]string{
	"mysql":      "LIKE",
	"postgresql": "ILIKE",
}

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) Migration() error {
	return r.db.AutoMigrate(&model.Book{})
}

func (r *BookRepository) FindAll(ctx context.Context, includeDeleted bool) ([]model.Book, error) {
	var books []model.Book
	var result *gorm.DB

	tx := r.db.WithContext(ctx)
	if includeDeleted {
		result = tx.Unscoped().Find(&books)
	} else {
		result = tx.Find(&books)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (r *BookRepository) FindById(ctx context.Context, id uint) (model.Book, error) {
	var b model.Book
	tx := r.db.WithContext(ctx)
	result := tx.Preload("Author").First(&b, id)
	if result.Error != nil {
		return model.Book{}, result.Error
	}
	return b, nil
}

func (r *BookRepository) Search(ctx context.Context, text string) ([]model.Book, error) {
	var books []model.Book
	text = "%" + text + "%"
	likeOp, ok := LikeOperator[os.Getenv("PATIKA_DB_DRIVER")]
	if !ok {
		return nil, errors.New("invalid db driver")
	}

	queryString := fmt.Sprintf("bakyazi_books.Name %s ? OR bakyazi_books.ISBN %s ? OR aut.Name %s ?",
		likeOp, likeOp, likeOp)
	tx := r.db.WithContext(ctx)
	result := tx.Preload("Author").
		Joins("join bakyazi_authors aut on aut.id = bakyazi_books.author_id").
		Where(queryString, text, text, text).
		Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (r *BookRepository) Update(ctx context.Context, b *model.Book) error {
	tx := r.db.WithContext(ctx)
	result := tx.Save(b)
	return result.Error
}

func (r *BookRepository) Delete(ctx context.Context, b model.Book) error {
	tx := r.db.WithContext(ctx)
	result := tx.Delete(&b)
	return result.Error
}

func (r *BookRepository) DeleteByAuthorId(ctx context.Context, id uint) error {
	tx := r.db.WithContext(ctx)
	return tx.Where("author_id = ?", id).Delete(&model.Book{}).Error
}

func (r *BookRepository) Insert(ctx context.Context, book *model.Book) error {
	return r.db.WithContext(ctx).Preload("Author").Create(&book).Error
}

func (r *BookRepository) Clear(ctx context.Context) error {
	return r.db.WithContext(ctx).Unscoped().Where("id > ?", 0).Delete(&model.Book{}).Error
}

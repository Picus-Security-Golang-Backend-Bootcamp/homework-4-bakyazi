package repository

import (
	"context"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/db/domain/model"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (r *AuthorRepository) Migration() error {
	return r.db.AutoMigrate(&model.Author{})
}

func (r *AuthorRepository) FindAll(ctx context.Context, includeDeleted bool) ([]model.Author, error) {
	var authors []model.Author
	var result *gorm.DB

	tx := r.db.WithContext(ctx)
	if includeDeleted {
		result = tx.Unscoped().Find(&authors)
	} else {
		result = tx.Find(&authors)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}

func (r *AuthorRepository) FindById(ctx context.Context, id uint) (model.Author, error) {
	var a model.Author
	tx := r.db.WithContext(ctx)
	result := tx.Preload("Books").First(&a, id)
	if result.Error != nil {
		return model.Author{}, result.Error
	}
	return a, nil
}

func (r *AuthorRepository) Insert(ctx context.Context, author *model.Author) error {
	result := r.db.WithContext(ctx).Create(author)
	return result.Error
}

func (r *AuthorRepository) Update(ctx context.Context, author *model.Author) error {
	tx := r.db.WithContext(ctx)
	result := tx.Save(author)
	return result.Error
}

func (r *AuthorRepository) Delete(ctx context.Context, a model.Author) error {
	return r.db.WithContext(ctx).Delete(&a).Error
}

func (r *AuthorRepository) Clear(ctx context.Context) error {
	return r.db.WithContext(ctx).Unscoped().Where("id > ?", 0).Delete(&model.Author{}).Error
}

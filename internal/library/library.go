package library

import (
	"context"
	"errors"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/db/domain/model"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/db/domain/repository"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/db/infrastructure"
	"gorm.io/gorm"
	"os"
)

var (
	bookRepo   *repository.BookRepository
	authorRepo *repository.AuthorRepository
)

// Init initializes DB, repositories and insert sample data if DB is empty
func Init() {
	var db *gorm.DB
	switch os.Getenv("PATIKA_DB_DRIVER") {
	case "postgresql":
		db = infrastructure.NewPostgresDB()
	case "mysql":
		db = infrastructure.NewMySQLDB()
	default:
		panic("invalid PATIKA_DB_DRIVER")
	}

	bookRepo = repository.NewBookRepository(db)
	authorRepo = repository.NewAuthorRepository(db)

	err := authorRepo.Migration()
	if err != nil {
		panic(err)
	}
	err = bookRepo.Migration()
	if err != nil {
		panic(err)
	}

	insertSampleData()
}

// ListBooks service layer of list operation
// it returns all books not deleted
func ListBooks(ctx context.Context) ([]model.Book, error) {
	books, err := bookRepo.FindAll(ctx, false)
	if err != nil {
		return nil, err
	}
	return books, nil
}

// SearchBooks service layer of search operation
// it finds books not deleted and meet criteria
func SearchBooks(ctx context.Context, text string) ([]model.Book, error) {
	if len(text) < 3 {
		return nil, ErrTooShortSearchText
	}
	books, err := bookRepo.Search(ctx, text)
	if err != nil {
		return nil, err
	}
	return books, nil
}

// GetBook returns book by id
func GetBook(ctx context.Context, id uint) (*model.Book, error) {
	b, err := bookRepo.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBookNotFound
		}
		return nil, err
	}
	return &b, nil
}

// UpdateBook update book by id
func UpdateBook(ctx context.Context, b *model.Book, id uint) (*model.Book, error) {
	if id == 0 {
		return nil, ErrInvalidBookId
	}

	if id > 0 {
		_, err := bookRepo.FindById(ctx, id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrBookNotFound
			}
			return nil, err
		}
	}
	b.ID = id
	err := bookRepo.Update(ctx, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// BuyBook service layer of buy operation
// it decreases stock amount of book with given id by quantity
func BuyBook(ctx context.Context, id uint, quantity int) (*model.Book, error) {

	b, err := bookRepo.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBookNotFound
		}
		return nil, err
	}
	if b.StockAmount >= quantity {
		b.StockAmount -= quantity
		err = bookRepo.Update(ctx, &b)
		if err != nil {
			return nil, err
		}
		return &b, nil
	}
	return nil, ErrBookOutOfStock
}

// DeleteBook service layer of delete operation
// deletes book with given id if it is nor already deleted
func DeleteBook(ctx context.Context, id uint) (*model.Book, error) {
	b, err := bookRepo.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBookNotFound
		}
		return nil, err
	}
	err = bookRepo.Delete(ctx, b)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

// CreateBook creates new Book
func CreateBook(ctx context.Context, b model.Book) (*model.Book, error) {
	if b.ID > 0 {
		_, err := bookRepo.FindById(ctx, b.ID)
		if err == nil {
			return nil, ErrBookIdAlreadyExist
		}
	}
	err := bookRepo.Insert(ctx, &b)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

// GetAuthor returns author by id
func GetAuthor(ctx context.Context, id uint) (*model.Author, error) {
	a, err := authorRepo.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAuthorNotFound
		}
		return nil, err
	}
	return &a, nil
}

// ListAuthors return list of all authors
func ListAuthors(ctx context.Context) ([]model.Author, error) {
	result, err := authorRepo.FindAll(ctx, false)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAuthor creates new Author
func CreateAuthor(ctx context.Context, a model.Author) (*model.Author, error) {
	if a.ID > 0 {
		_, err := authorRepo.FindById(ctx, a.ID)
		if err == nil {
			return nil, ErrAuthorIdAlreadyExist
		}
	}
	err := authorRepo.Insert(ctx, &a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// UpdateAuthor updates Author by id
func UpdateAuthor(ctx context.Context, a *model.Author, id uint) (*model.Author, error) {
	if id == 0 {
		return nil, ErrInvalidAuthorId
	}

	if id > 0 {
		_, err := authorRepo.FindById(ctx, id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrAuthorNotFound
			}
			return nil, err
		}
	}
	a.ID = id
	err := authorRepo.Update(ctx, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// DeleteAuthor deletes author
func DeleteAuthor(ctx context.Context, id uint) error {
	a, err := authorRepo.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAuthorNotFound
		}
		return err
	}

	err = bookRepo.DeleteByAuthorId(ctx, id)
	if err != nil {
		return err
	}

	return authorRepo.Delete(ctx, a)
}

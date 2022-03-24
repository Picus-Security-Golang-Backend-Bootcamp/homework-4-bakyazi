package library

import (
	"errors"
	"net/http"
)

type ApiError struct {
	Status int
	error
}

var (
	ErrBookOutOfStock     = ApiError{http.StatusNotAcceptable, errors.New("there is not enough stock to sell this book in demanded amount")}
	ErrBookNotFound       = ApiError{http.StatusNotFound, errors.New("book not found")}
	ErrInvalidBookId      = ApiError{http.StatusBadRequest, errors.New("book id should be positive numbers")}
	ErrBookIdAlreadyExist = ApiError{http.StatusBadRequest, errors.New("book id is already used")}
	ErrTooShortSearchText = ApiError{http.StatusBadRequest, errors.New("search text is too short, should be greater than 2")}
)

var (
	ErrInvalidAuthorId      = ApiError{http.StatusBadRequest, errors.New("book id should be positive numbers")}
	ErrAuthorNotFound       = ApiError{http.StatusNotFound, errors.New("author not found")}
	ErrAuthorIdAlreadyExist = ApiError{http.StatusBadRequest, errors.New("author id is already used")}
)

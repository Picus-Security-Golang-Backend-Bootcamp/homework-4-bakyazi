package books

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/api/helper"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/api/response"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/db/domain/model"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/library"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// BuyRequest form body of buy request
type BuyRequest struct {
	BookID   uint `json:"book_id"`
	Quantity int  `json:"quantity"`
}

// SetRouters sets books endpoint handlers
func SetRouters(r *mux.Router) {
	b := r.PathPrefix("/books").Subrouter()
	b.HandleFunc("/{id:[0-9]+}", bookIdHandler).Methods(
		http.MethodGet,
		http.MethodPut,
		http.MethodDelete,
	)
	b.HandleFunc("/buy", buyBookHandler).Methods(http.MethodPost)
	b.HandleFunc("/search", searchBooks).Methods(http.MethodGet)
	b.HandleFunc("", bookHandler).Methods(
		http.MethodGet,
		http.MethodPost,
	)
}

// buyBookHandler buys an Book
func buyBookHandler(writer http.ResponseWriter, request *http.Request) {
	var br BuyRequest
	var err error
	writer.Header().Set("Content-Type", "application/json")

	err = helper.DecodeJSONBody(writer, request, &br)
	if err != nil {
		response.WriteResponse(writer, response.NewErrorResponse(err))
		return
	}
	_, err = library.BuyBook(request.Context(), br.BookID, br.Quantity)
	if err != nil {
		response.WriteResponse(writer, response.NewErrorResponse(err))
		return
	}
	response.WriteResponse(writer, response.NewSuccessResponse(http.StatusOK, "Success."))
}

// bookIdHandler switches handlers depends on method for /books/{id} endpoint
func bookIdHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		getBook(writer, request)
	case http.MethodPut:
		putBook(writer, request)
	case http.MethodDelete:
		deleteBook(writer, request)
	}
}

// deleteBook deletes a Book
func deleteBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	writer.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(vars["id"])
	if id <= 0 {
		response.WriteResponse(writer, response.NewErrorResponse(library.ErrInvalidBookId))
		return
	}
	_, err := library.DeleteBook(request.Context(), uint(id))
	if err != nil {
		response.WriteResponse(writer, response.NewErrorResponse(err))
		return
	}
	response.WriteResponse(writer, response.NewSuccessResponse(http.StatusOK, "Success."))
}

//putBook updates a book
func putBook(writer http.ResponseWriter, request *http.Request) {
	var b model.Book
	var err error
	vars := mux.Vars(request)
	writer.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(vars["id"])
	if id <= 0 {
		response.WriteResponse(writer, response.NewErrorResponse(library.ErrInvalidBookId))
		return
	}

	err = helper.DecodeJSONBody(writer, request, &b)
	if err != nil {
		response.WriteResponse(writer, response.NewErrorResponse(err))
		return
	}
	updatedBook, err := library.UpdateBook(request.Context(), &b, uint(id))
	if err != nil {
		response.WriteResponse(writer, response.NewErrorResponse(err))
		return
	}
	response.WriteResponse(writer, response.NewSuccessResponse(http.StatusOK, updatedBook))
}

// getBook returns a book
func getBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	writer.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(vars["id"])
	if id <= 0 {
		response.WriteResponse(writer, response.NewErrorResponse(library.ErrInvalidBookId))
		return
	}
	_book, err := library.GetBook(request.Context(), uint(id))
	if err != nil {
		response.WriteResponse(writer, response.NewErrorResponse(err))
		return
	}
	response.WriteResponse(writer, response.NewSuccessResponse(http.StatusOK, _book))
}

// bookHandler switches handlers depends on method for /books endpoint
func bookHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		getBooks(writer, request)
	case http.MethodPost:
		postBook(writer, request)
	}
}

// postBook creates new book
func postBook(writer http.ResponseWriter, request *http.Request) {
	var b model.Book
	writer.Header().Set("Content-Type", "application/json")

	err := helper.DecodeJSONBody(writer, request, &b)
	if err != nil {
		response.WriteResponse(writer, response.NewErrorResponse(err))
		return
	}
	createdBook, err := library.CreateBook(request.Context(), b)
	if err != nil {
		response.WriteResponse(writer, response.NewErrorResponse(err))
		return
	}
	response.WriteResponse(writer, response.NewSuccessResponse(http.StatusOK, createdBook))
}

// getBooks return all books
func getBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	books, err := library.ListBooks(request.Context())
	if err != nil {
		response.WriteResponse(writer, response.NewErrorResponse(err))
		return
	}
	response.WriteResponse(writer, response.NewSuccessResponse(http.StatusOK, books))
}

// searchBooks return list of books matching searched pattern
func searchBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	query := request.URL.Query().Get("query")
	books, err := library.SearchBooks(request.Context(), query)
	if err != nil {
		response.WriteResponse(writer, response.NewErrorResponse(err))
		return
	}
	response.WriteResponse(writer, response.NewSuccessResponse(http.StatusOK, books))
}

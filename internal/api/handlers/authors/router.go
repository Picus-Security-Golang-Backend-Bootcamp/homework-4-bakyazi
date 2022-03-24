package authors

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/api/helper"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/api/response"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/db/domain/model"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/library"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// SetRouters sets author endpoint handlers
func SetRouters(r *mux.Router) {
	a := r.PathPrefix("/authors").Subrouter()
	a.HandleFunc("/{id:[0-9]+}", authorIdHandler).Methods(
		http.MethodGet,
		http.MethodPut,
		http.MethodDelete,
	)
	a.HandleFunc("", authorHandler).Methods(
		http.MethodGet,
		http.MethodPost,
	)
}

// authorHandler switches handlers depends on method for /authors endpoint
func authorHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		getAuthors(writer, request)
	case http.MethodPost:
		postAuthor(writer, request)
	}
}

// authorIdHandler switches handlers depends on method for /authors/{id} endpoint
func authorIdHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		getAuthor(writer, request)
	case http.MethodPut:
		putAuthor(writer, request)
	case http.MethodDelete:
		deleteAuthor(writer, request)
	}
}

// getAuthors returns all authors
func getAuthors(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	out, err := library.ListAuthors(request.Context())
	if err != nil {
		response.WriteResponse(writer, response.NewErrorResponse(err))
		return
	}
	response.WriteResponse(writer, response.NewSuccessResponse(http.StatusOK, out))
}

// postAuthor creates new author
func postAuthor(writer http.ResponseWriter, request *http.Request) {
	var a model.Author
	writer.Header().Set("Content-Type", "application/json")

	err := helper.DecodeJSONBody(writer, request, &a)
	if err != nil {
		response.WriteResponse(writer, response.NewErrorResponse(err))
		return
	}
	createdAuthor, err := library.CreateAuthor(request.Context(), a)
	if err != nil {
		response.WriteResponse(writer, response.NewErrorResponse(err))
		return
	}
	response.WriteResponse(writer, response.NewSuccessResponse(http.StatusOK, createdAuthor))
}

// putAuthor updates an Author
func putAuthor(writer http.ResponseWriter, request *http.Request) {
	var a model.Author
	var err error
	vars := mux.Vars(request)
	writer.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(vars["id"])
	if id <= 0 {
		response.WriteResponse(writer, response.NewErrorResponse(library.ErrInvalidAuthorId))
		return
	}
	err = helper.DecodeJSONBody(writer, request, &a)
	if err != nil {
		response.WriteResponse(writer, response.NewErrorResponse(err))
		return
	}
	updatedAuthor, err := library.UpdateAuthor(request.Context(), &a, uint(id))
	if err != nil {
		response.WriteResponse(writer, response.NewErrorResponse(err))
		return
	}
	response.WriteResponse(writer, response.NewSuccessResponse(http.StatusOK, updatedAuthor))
}

// getAuthor return an Author by id
func getAuthor(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	writer.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(vars["id"])
	if id <= 0 {
		response.WriteResponse(writer, response.NewErrorResponse(library.ErrInvalidAuthorId))
		return
	}
	_author, err := library.GetAuthor(request.Context(), uint(id))
	if err != nil {
		response.WriteResponse(writer, response.NewErrorResponse(err))
		return
	}
	response.WriteResponse(writer, response.NewSuccessResponse(http.StatusOK, _author))

}

// deleteAuthor deletes an Author
func deleteAuthor(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	writer.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(vars["id"])
	if id <= 0 {
		response.WriteResponse(writer, response.NewErrorResponse(library.ErrInvalidAuthorId))
		return
	}
	err := library.DeleteAuthor(request.Context(), uint(id))
	if err != nil {
		response.WriteResponse(writer, response.NewErrorResponse(err))
		return
	}
	response.WriteResponse(writer, response.NewSuccessResponse(http.StatusOK, "Success."))

}

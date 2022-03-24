package response

import (
	"encoding/json"
	"fmt"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/library"
	"net/http"
)

type Response interface {
	Respond(http.ResponseWriter)
}

type ErrorResponse struct {
	StatusCode int
	Err        error
}

func (a ErrorResponse) Respond(writer http.ResponseWriter) {
	writer.WriteHeader(a.StatusCode)
	writer.Write([]byte(a.Err.Error()))
}

type SuccessResponse struct {
	StatusCode int
	Data       []byte
}

func (s SuccessResponse) Respond(writer http.ResponseWriter) {
	writer.WriteHeader(s.StatusCode)
	writer.Write(s.Data)
}

func WriteResponse(w http.ResponseWriter, r Response) {
	r.Respond(w)
}

func NewErrorResponse(err error) Response {
	switch err.(type) {
	case library.ApiError:
		fmt.Println("dsadsadasdas")
		return ErrorResponse{
			StatusCode: err.(library.ApiError).Status,
			Err:        err,
		}
	default:
		return ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
}

func NewSuccessResponse(status int, data interface{}) Response {
	d, err := json.Marshal(data)
	if err != nil {
		return NewErrorResponse(err)
	}
	return SuccessResponse{
		StatusCode: status,
		Data:       d,
	}
}

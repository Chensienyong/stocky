package customerror

import (
	"net/http"

	"github.com/chensienyong/stocky/pkg/response"
)

//New custom error
func New(httpStatusCode, code int, message, field string) response.CustomError {
	return response.NewErr(httpStatusCode, code, field, message)
}

var (
	RecordNotFound = New(http.StatusNotFound, 404, "Record Not Found", "")
	DBError = New(http.StatusInternalServerError, 500, "Issue with Database", "")
	RedisError = New(http.StatusInternalServerError, 500, "Issue with Redis", "")
)

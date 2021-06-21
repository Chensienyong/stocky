package response_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/chensienyong/stocky/pkg/response"
	"github.com/stretchr/testify/assert"
)

func TestCustomError_Error(t *testing.T) {
	ce := response.CustomError{Message: "stocky_error"}
	assert.Equal(t, ce.Message, ce.Error())
}

func TestBuildSuccess(t *testing.T) {
	data    := "data"
	message := "request accepted"
	meta    := "meta"

	res := response.BuildSuccess(data, message, meta)
	assert.NotNil(t, res)
	assert.Equal(t, data, res.Data)
	assert.Equal(t, message, res.Message)
	assert.Equal(t, meta, res.Meta)
}

func TestBuildError(t *testing.T) {
	res := response.BuildError()
	assert.NotNil(t, res)
	assert.Equal(t, res.Errors[0].Code, response.UnexpectedErr.Code)

	res = response.BuildError(errors.New("Error"))
	assert.NotNil(t, res)

	ce := response.CustomError{}
	res = response.BuildError(ce)
	assert.NotNil(t, res)
}

type Writer struct{}

func (w Writer) Header() http.Header {
	return make(http.Header)
}

func (w Writer) Write(b []byte) (int, error) {
	return 0, nil
}

func (w Writer) WriteHeader(x int) {}

func TestWrite(t *testing.T) {
	writer := Writer{}
	assert.NotPanics(t, func() { response.Write(writer, "data", 200) })
}

func TestNewErr(t *testing.T) {
	httpCode := 200
	code     := 201
	field    := "field"
	message  := "message"

	error := response.NewErr(httpCode, code, field, message)

	assert.IsType(t, response.CustomError{}, error)
	assert.Equal(t, httpCode, error.HTTPCode)
}
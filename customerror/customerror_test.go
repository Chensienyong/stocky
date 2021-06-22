package customerror_test

import (
	"github.com/chensienyong/stocky/customerror"
	"github.com/chensienyong/stocky/pkg/response"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	error := customerror.New(400, 400, "message", "field")

	assert.IsType(t, response.CustomError{}, error)
	assert.Equal(t, 400, error.HTTPCode)
}

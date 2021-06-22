package handler_test

import (
	"context"
	"github.com/chensienyong/stocky/handler"
	"github.com/chensienyong/stocky/pkg/response"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOKWithIndex(t *testing.T) {
	writer := Writer{}
	assert.NotPanics(t, func() { handler.OKWithIndex(writer, nil, "", nil) })
}

func TestOK(t *testing.T) {
	writer := Writer{}
	assert.NotPanics(t, func() { handler.OK(writer, nil, "") })
}

func TestOKWithServerTime(t *testing.T) {
	writer := Writer{}
	assert.NotPanics(t, func() { handler.OKWithServerTime(writer, nil, "") })
}

func TestCreated(t *testing.T) {
	writer := Writer{}
	assert.NotPanics(t, func() { handler.Created(writer, nil) })
}

func TestCreatedWithMessage(t *testing.T) {
	writer := Writer{}
	assert.NotPanics(t, func() { handler.CreatedWithMessage(writer, nil, "") })
}

func TestAccepted(t *testing.T) {
	writer := Writer{}
	assert.NotPanics(t, func() { handler.Accepted(writer, nil, "") })
}

func TestErrorInfo(t *testing.T) {
	writer := Writer{}
	assert.NotPanics(t, func() { handler.ErrorInfo(writer, nil) })
}

func TestError(t *testing.T) {
	writer := Writer{}
	assert.NotPanics(t, func() { handler.Error(writer, nil) })
}

func TestError_DeadlineExceeded(t *testing.T) {
	writer := Writer{}
	assert.NotPanics(t, func() { handler.Error(writer, context.DeadlineExceeded) })
}

func TestError_CustomError(t *testing.T) {
	writer := Writer{}
	assert.NotPanics(t, func() { handler.Error(writer, response.UnexpectedErr) })
}

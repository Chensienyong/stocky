package handler_test

import (
	"github.com/chensienyong/stocky"
	"github.com/chensienyong/stocky/handler"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewHandler(t *testing.T) {
	newHandler := handler.NewHandler(&stocky.Stocky{})

	assert.IsType(t, &handler.Handler{}, newHandler)
}

type Writer struct{
	status int
}

func (w Writer) Header() http.Header {
	return make(http.Header)
}

func (w Writer) Write(b []byte) (int, error) {
	return 0, nil
}

func (w Writer) WriteHeader(x int) {
	w.status = x
}

func TestHandler_Healthz(t *testing.T) {
	newHandler := handler.Handler{}
	writer := Writer{}
	assert.NotPanics(t, func() { newHandler.Healthz(writer, nil, nil) })
}

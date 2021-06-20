package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/chensienyong/stocky/pkg/response"
	"github.com/pkg/errors"
)

// IndexMeta contains metadata of an index response
type IndexMeta struct {
	HTTPStatus int    `json:"http_status"`
	Limit      uint64 `json:"limit"`
	Offset     uint64 `json:"offset"`
	Total      uint64 `json:"total"`
}

// ServerMeta contains metadata of an server time response
type ServerMeta struct {
	HTTPStatus int    `json:"http_status"`
	ServerTime string `json:"server_time,omitempty"`
}

// OKWithIndex wrap success response with meta offset
func OKWithIndex(w http.ResponseWriter, data interface{}, message string, meta interface{}) {
	successResponse := response.BuildSuccess(data, message, meta)
	response.Write(w, successResponse, http.StatusOK)
}

// OK wrap success response with offset
func OK(w http.ResponseWriter, data interface{}, message string) {
	successResponse := response.BuildSuccess(data, message, response.MetaInfo{HTTPStatus: http.StatusOK})
	response.Write(w, successResponse, http.StatusOK)
}

// OKWithServerTime wrap success response with server time
func OKWithServerTime(w http.ResponseWriter, data interface{}, message string) {
	serverTime := time.Now().Format(time.RFC3339)
	successResponse := response.BuildSuccess(data, message, ServerMeta{HTTPStatus: http.StatusOK, ServerTime: serverTime})
	response.Write(w, successResponse, http.StatusOK)
}

// Created wrap create response
func Created(w http.ResponseWriter, data interface{}) {
	successResponse := response.BuildSuccess(data, "Created", response.MetaInfo{HTTPStatus: http.StatusCreated})
	response.Write(w, successResponse, http.StatusCreated)
}

// CreatedWithMessage wrap create response with message
func CreatedWithMessage(w http.ResponseWriter, data interface{}, message string) {
	successResponse := response.BuildSuccess(data, message, response.MetaInfo{HTTPStatus: http.StatusCreated})
	response.Write(w, successResponse, http.StatusCreated)
}

// Accepted wrap update response with message
func Accepted(w http.ResponseWriter, data interface{}, message string) {
	successResponse := response.BuildSuccess(data, message, response.MetaInfo{HTTPStatus: http.StatusAccepted})
	response.Write(w, successResponse, http.StatusAccepted)
}

// ErrorInfo wrap error info response
func ErrorInfo(w http.ResponseWriter, errors []error) {
	errorResponse := response.BuildError(errors...)
	response.Write(w, errorResponse, http.StatusUnprocessableEntity)
	return
}

// Error wrap error response
func Error(w http.ResponseWriter, err error) {
	if err == context.Canceled || err == context.DeadlineExceeded {
		return
	}

	if ce, ok := err.(response.CustomError); ok {
		errorResponse := response.BuildError(ce)
		response.Write(w, errorResponse, ce.HTTPCode)
		return
	}

	causer := errors.Cause(err)
	customError := causer

	errorResponse := response.BuildError(customError)
	meta := errorResponse.Meta.(response.MetaInfo)
	response.Write(w, errorResponse, meta.HTTPStatus)
}

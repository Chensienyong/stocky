package middleware

import (
	"github.com/chensienyong/stocky/pkg/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HandleWithError is an httprouter.Handle that returns an error.
type HandleWithError func(http.ResponseWriter, *http.Request, httprouter.Params) error

// Decorator decorates HandleWithError.
type Decorator func(HandleWithError) HandleWithError

// HTTP runs HandleWithError and converts it to httprouter.Handle.
// The conversion is needed because httprouter.Router needs httprouter.Handle
// in its signature.
func HTTP(handle HandleWithError) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		handle(w, r, params)
	}
}

// ApplyDecorators returns decorated HandleWithError.
func ApplyDecorators(handle HandleWithError, ds ...Decorator) HandleWithError {
	for _, d := range ds {
		handle = d(handle)
	}
	return handle
}

// WithBasicAuth decorates Decorator with basic authorization.
func WithBasicAuth(requiredUser, requiredPassword string) Decorator {
	return func(handle HandleWithError) HandleWithError {
		return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
			// Get the Basic Authentication credentials
			user, password, hasAuth := r.BasicAuth()
			if hasAuth && user == requiredUser && password == requiredPassword {
				// Delegate request to the given handle
				return handle(w, r, params)
			}
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			err := response.CustomError{
				HTTPCode: http.StatusUnauthorized,
				Message:  "Unauthorized",
				Code:     401,
			}
			return err
		}
	}
}

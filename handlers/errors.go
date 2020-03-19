package handlers

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

type Error struct {
	err  error
	code int
}

func (e Error) Error() string {
	return e.err.Error()
}

var (
	NewError = func(e error) Error {
		return Error{
			err:  e,
			code: http.StatusInternalServerError,
		}
	}
	NotFound = func(err error) Error {
		return Error{
			err:  errors.Wrap(err, "not found"),
			code: http.StatusNotFound,
		}
	}
	Unauthorized = func(e error) Error {
		return Error{
			err:  e,
			code: http.StatusUnauthorized,
		}
	}
)

func RespondError(w http.ResponseWriter, err Error) {
	output := apiResponse{
		Status: "error",
		Error:  err.Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.code)
	json.NewEncoder(w).Encode(output)
}

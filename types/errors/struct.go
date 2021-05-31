package errors

import (
	"github.com/pkg/errors"
)

type appError struct {
	Err      string `json:"error"`
	Code     int    `json:"-"`
	DbCode   int    `json:"code,omitempty"`
	Id       int64  `json:"id,omitempty"`
	loggable bool   `json:"-"`
}

type Error interface {
	Error() string
	Status() int
}

func New(err string) Error {
	return &appError{
		Err: err,
	}
}

func Err(err Error) Error {
	return &appError{
		Err:  err.Error(),
		Code: err.Status(),
	}
}

func Wrap(err error, message string) Error {
	return &appError{
		Err: errors.Wrap(err, message).Error(),
	}
}

func (e *appError) Error() string {
	return e.Err
}

func (e appError) Status() int {
	if e.Code == 0 {
		return 200
	}
	return e.Code
}

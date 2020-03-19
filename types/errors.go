package types

import (
	"github.com/pkg/errors"
	"net/http"
)

var (
	ErrorServiceSelection = returnErr("error selecting services")

	// create errors
	ErrorCreateService    = returnErr("error creating service")
	ErrorCreateMessage    = returnErr("error creating messages")
	ErrorCreateIncident   = returnErr("error creating incident")
	ErrorCreateUser       = returnErr("error creating user")
	ErrorCreateIncidentUp = returnErr("error creating incident update")
	ErrorCreateGroup      = returnErr("error creating group")
	ErrorCreateCheckinHit = returnErr("error creating checkin hit")
	ErrorCreateSampleHits = returnErr("error creating sample hits")
	ErrorCreateCore       = returnErr("error creating core")
	ErrorCreateHit        = returnErr("error creating hit for service %v")

	ErrorDirCreate = returnErr("error creating directory %s")

	ErrorFileCopy = returnErr("error copying file %s to %s")

	ErrorConfig     = returnErr("error with configuration")
	ErrorConnection = returnErr("error with connection")

	ErrorNotFound  = returnErrCode("item was not found", http.StatusNotFound)
	ErrorJSONParse = returnErrCode("could not parse JSON request", http.StatusBadRequest)
)

type Errorer interface {
}

type Error struct {
	err  error
	code int
}

func (e Error) Error() string {
	return e.err.Error()
}

func (e Error) String() string {
	return e.err.Error()
}

func returnErrCode(str string, code int) error {
	return Error{
		err:  errors.New(str),
		code: code,
	}
}

func returnErr(str string) Error {
	return Error{
		err: errors.New(str),
	}
}

func convertError(val interface{}) string {
	switch v := val.(type) {
	case *Error:
		return v.Error()
	case string:
		return v
	default:
		return ""
	}
}

type errorer interface {
	Error() string
}

func ErrWrap(err errorer, format interface{}, args ...interface{}) Error {
	return Error{
		err:  errors.Wrapf(err, convertError(format), args...),
		code: 0,
	}
}

func Err(err errorer, format interface{}) Error {
	return Error{
		err:  errors.Wrap(err, convertError(format)),
		code: 0,
	}
}

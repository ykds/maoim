package merror

import (
	"fmt"
	perr "github.com/pkg/errors"
)

var errorMap = make(map[int32]*Error)


type Error struct {
	Code int32 `json:"code"`
	Message string `json:"message"`
}

func New(code int32, message string) *Error {
	if _, ok := errorMap[code]; ok {
		panic(fmt.Sprintf("错误码(%d)已被占用", code))
	}
	e := &Error{
		Code: code,
		Message: message,
	}
	errorMap[code] = e
	return e
}

func (e *Error) Error() string {
	return e.Message
}

func Wrap(err error, msg string) error {
	return perr.Wrap(err, msg)
}

func Unwrap(err error) error {
	return perr.Unwrap(err)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return perr.Wrapf(err, format, args...)
}

func As(err error, target interface{}) bool  {
	return perr.As(err, target)
}

func Is(err error, target error) bool {
	return perr.Is(err, target)
}

func WithMessage(err error, message string) error {
	return perr.WithMessage(err, message)
}

func WithMessagef(err error, format string, args ...interface{}) error {
	return perr.WithMessagef(err, format, args...)
}
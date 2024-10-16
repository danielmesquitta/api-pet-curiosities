package errs

import (
	"fmt"
	"runtime/debug"
)

type Err struct {
	Message    string `json:"message,omitempty"`
	StackTrace string `json:"stack_trace,omitempty"`
	Type       Type   `json:"type,omitempty"`
}

type Type string

const (
	ErrTypeUnknown      Type = "unknown"
	ErrTypeNotFound     Type = "not_found"
	ErrTypeUnauthorized Type = "unauthorized"
	ErrTypeForbidden    Type = "forbidden"
	ErrTypeValidation   Type = "validation_error"
)

func newErr(err any, t Type) *Err {
	switch v := err.(type) {
	case *Err:
		return v
	case error:
		return &Err{
			Message:    v.Error(),
			StackTrace: string(debug.Stack()),
			Type:       t,
		}
	case string:
		return &Err{
			Message:    v,
			StackTrace: string(debug.Stack()),
			Type:       t,
		}
	default:
		panic("trying to create an Err with an unsupported type")
	}
}

// NewErr creates a new Err instance from either an error or a string,
// and sets the Type flag to unknown. This is useful when you want to
// create an error that is not expected to happen, and you want to
// log it with stack tracing.
func New(err any) *Err {
	return newErr(err, ErrTypeUnknown)
}

func (e *Err) Error() string {
	return e.Message
}

func (e *Err) ErrorWithStackTrace() string {
	return fmt.Sprintf("%s\n\n%s", e.Message, e.StackTrace)
}

var (
	ErrUserNotFound = newErr(
		"user not found",
		ErrTypeNotFound,
	)
	ErrValidation = newErr(
		"validation error",
		ErrTypeValidation,
	)
	ErrUserNotAllowed = newErr(
		"only users with the PRO tier can execute this action",
		ErrTypeForbidden,
	)
	ErrPetNotFound = newErr(
		"pet not found",
		ErrTypeNotFound,
	)
)

var _ error = (*Err)(nil)

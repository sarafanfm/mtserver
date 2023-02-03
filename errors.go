package mtserver

import (
	"errors"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	err_ok        = "ok"
	err_canceled  = "cancelled"
	err_unknown   = "unknown"
	err_invalid   = "invalid argument"
	err_deadline  = "deadline exceeded"
	err_notfound  = "not found"
	err_exists    = "already exists"
	err_forbid    = "permission denied"
	err_exhausted = "resource exhausted"
	err_precond   = "failed precondition"
	err_aborted   = "aborted"
	err_range     = "out of range"
	err_unimpl    = "unimplemented"
	err_internal  = "internal"
	err_unavail   = "unavailable"
	err_data      = "data loss"
	err_unauth    = "unauthenticated"
)

type Exception error

var (
	ErrOK       Exception = errors.New(err_ok)
	ErrCanceled Exception = errors.New(err_canceled)
	ErrUnknown  Exception = errors.New(err_unknown)
	ErrInvalid  Exception = errors.New(err_invalid)
	ErrDeadline Exception = errors.New(err_deadline)
	ErrNotFound Exception = errors.New(err_notfound)
	ErrExists   Exception = errors.New(err_exists)
	ErrForbid   Exception = errors.New(err_forbid)
	ErrExhaust  Exception = errors.New(err_exhausted)
	ErrPrecond  Exception = errors.New(err_precond)
	ErrAborted  Exception = errors.New(err_aborted)
	ErrRange    Exception = errors.New(err_range)
	ErrUnimpl   Exception = errors.New(err_unimpl)
	ErrInternal Exception = errors.New(err_internal)
	ErrUnavail  Exception = errors.New(err_unavail)
	ErrData     Exception = errors.New(err_data)
	ErrUnauth   Exception = errors.New(err_unauth)
)

var grpcErrors = map[error]codes.Code{
	ErrOK:       codes.OK,
	ErrCanceled: codes.Canceled,
	ErrUnknown:  codes.Unknown,
	ErrInvalid:  codes.InvalidArgument,
	ErrDeadline: codes.DeadlineExceeded,
	ErrNotFound: codes.NotFound,
	ErrExists:   codes.AlreadyExists,
	ErrForbid:   codes.PermissionDenied,
	ErrExhaust:  codes.ResourceExhausted,
	ErrPrecond:  codes.FailedPrecondition,
	ErrAborted:  codes.Aborted,
	ErrRange:    codes.OutOfRange,
	ErrUnimpl:   codes.Unimplemented,
	ErrInternal: codes.Internal,
	ErrUnavail:  codes.Unavailable,
	ErrData:     codes.DataLoss,
	ErrUnauth:   codes.Unauthenticated,
}

var httpErrors = map[error]int{
	ErrOK:       http.StatusOK,
	ErrCanceled: http.StatusRequestTimeout,
	ErrUnknown:  http.StatusInternalServerError,
	ErrInvalid:  http.StatusBadRequest,
	ErrDeadline: http.StatusRequestTimeout,
	ErrNotFound: http.StatusNotFound,
	ErrExists:   http.StatusConflict,
	ErrForbid:   http.StatusForbidden,
	ErrExhaust:  http.StatusTooManyRequests,
	ErrPrecond:  http.StatusPreconditionFailed,
	ErrAborted:  http.StatusConflict,
	ErrRange:    http.StatusRequestedRangeNotSatisfiable,
	ErrUnimpl:   http.StatusNotImplemented,
	ErrInternal: http.StatusInternalServerError,
	ErrUnavail:  http.StatusServiceUnavailable,
	ErrData:     http.StatusInsufficientStorage,
	ErrUnauth:   http.StatusUnauthorized,
}

func NewError(message string, errType Exception) error {
	return fmt.Errorf("%w: %s", errType, message)
}

func GrpcError(err error) error {
	if err == nil {
		return nil
	}
	if grpcCode, ok := grpcErrors[errors.Unwrap(err)]; ok {
		return status.Error(grpcCode, err.Error())
	}
	return err
}

func HttpCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	if httpCode, ok := httpErrors[errors.Unwrap(err)]; ok {
		return httpCode
	}
	return http.StatusInternalServerError
}

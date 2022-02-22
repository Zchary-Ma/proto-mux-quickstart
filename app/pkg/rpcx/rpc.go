package rpcx

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HTTPStatusFromCode converts a gRPC error code into the corresponding HTTP response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func HTTPStatusFromCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.InvalidArgument, codes.FailedPrecondition, codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists, codes.Aborted:
		return http.StatusConflict
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.Canceled:
		return 499
	case codes.Unknown, codes.Internal, codes.DataLoss:
		return http.StatusInternalServerError
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	}
	return http.StatusInternalServerError
}

// BadRequest indicates client specified an invalid argument.
func BadRequest(format string, a ...interface{}) error {
	return status.Errorf(codes.InvalidArgument, format, a...)
}

// NotFound means some requested entity (e.g., file or directory) was
// not found.
func NotFound(format string, a ...interface{}) error {
	return status.Errorf(codes.NotFound, format, a...)
}

// AlreadyExists means an attempt to create an entity failed because one
// already exists.
func AlreadyExists(format string, a ...interface{}) error {
	return status.Errorf(codes.AlreadyExists, format, a...)
}

// PermissionDenied indicates the caller does not have permission to
// execute the specified operation.
func PermissionDenied(format string, a ...interface{}) error {
	return status.Errorf(codes.PermissionDenied, format, a...)
}

// Aborted indicates the operation was aborted, typically due to a
// concurrency issue like sequencer check failures, transaction aborts, etc.
func Aborted(format string, a ...interface{}) error {
	return status.Errorf(codes.Aborted, format, a...)
}

// OutOfRange means operation was attempted past the valid range.
// E.g., seeking or reading past end of file.
func OutOfRange(format string, a ...interface{}) error {
	return status.Errorf(codes.OutOfRange, format, a...)
}

// Internal errors. Means some invariants expected by underlying
// system has been broken. If you see one of these errors,
// something is very broken.
func Internal(format string, a ...interface{}) error {
	return status.Errorf(codes.Internal, format, a...)
}

// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 为某个枚举单独设置错误码
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NOT_FOUND.String() && e.Code == 404
}

// 为某个枚举单独设置错误码
func ErrorNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsBadRequest(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_BAD_REQUEST.String() && e.Code == 400
}

func ErrorBadRequest(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_BAD_REQUEST.String(), fmt.Sprintf(format, args...))
}

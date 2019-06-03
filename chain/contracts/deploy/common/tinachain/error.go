package tinachain

import (
	"fmt"
)

const (
	ErrorOK = iota

	ErrorParamNotFound = iota + 1000
	ErrorParamInvalid
	ErrorServerError
	ErrorJsonParseError
	ErrorRedisError
	ErrorMysqlError
	ErrorAppIdNotFound
	ErrorSignatureCheckFail
	ErrorFileExist
	ErrorBokerchainNotInitialized
	ErrorContractCallError
	ErrorAccountExist
	ErrorAccountNotFound
	ErrorAccountBinded
	ErrorNotAuthorized
	ErrorAppNotFound
	ErrorShaNotMatch
	ErrorNotMatch
)

type ErrorContext struct {
	Code    int
	Message string
}
type Error *ErrorContext

func NewError(code int, format string, args ...interface{}) Error {
	return &ErrorContext{code, fmt.Sprintf(format, args...)}
}

func NewServerError(format string, args ...interface{}) Error {
	return &ErrorContext{ErrorServerError, fmt.Sprintf(format, args...)}
}

func NewMysqlError(format string, args ...interface{}) Error {
	return &ErrorContext{ErrorMysqlError, fmt.Sprintf(format, args...)}
}

func NewRedisError(format string, args ...interface{}) Error {
	return &ErrorContext{ErrorRedisError, fmt.Sprintf(format, args...)}
}

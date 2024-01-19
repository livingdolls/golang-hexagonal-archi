package error_code

type ErrorCode string

const (
	Success        ErrorCode = "SUCCESS"
	InvalidRequest ErrorCode = "INVALID_REQUEST"
	DuplicateUser  ErrorCode = "DUPLICATE_USER"
	InternalError  ErrorCode = "INTERNAL_ERROR"
)

const (
	SuccessErrMsg        = "success"
	InternalErrMsg       = "internal error"
	InvalidRequestErrMsg = "invalid request"
)

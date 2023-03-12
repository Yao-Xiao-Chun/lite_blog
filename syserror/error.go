package syserror

/**
错误类型接口
*/
type Error interface {
	Code() int
	Error() string
	ReasonError() error
}

// New 调用方式
func New(msg string, reason error) Error {

	return UnknownError{msg, reason}

}

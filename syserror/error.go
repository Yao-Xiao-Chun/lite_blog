package syserror

/**
	错误类型接口
 */
type Error interface {
	Code() int
	Error() string
	ReasonError() error
}

/**
	返回值应该是Erro 单不明原因报错，返回interface

 */
func New(msg string,reson error) Error{

	return UnkownError{msg,reson}


}
package syserror

type UnknownError struct {
	msg string

	reason error
}

/**
code 值
*/
func (e UnknownError) Code() int { //不带指针

	return 1000
}

/**
实现error中的接口三个方法
*/
func (e UnknownError) Error() string {

	if len(e.msg) == 0 {

		return "未知错误"
	} else {

		return e.msg
	}
}

/**
直接返回错误原因

*/
func (e UnknownError) ReasonError() error {

	return e.reason
}

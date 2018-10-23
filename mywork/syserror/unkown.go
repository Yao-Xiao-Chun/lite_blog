package syserror

type UnkownError struct {

	msg string

	reson error
}

/**
	code 值
 */
func (this UnkownError)Code() int { //不带指针

	return 1000
}

/**
	实现error中的接口三个方法
 */
func (this UnkownError)Error() string{

	if len(this.msg) == 0{

		return "未知错误"
	}else{

		return this.msg
	}
}

/**
	直接返回错误原因

 */
func (this UnkownError)ReasonError() error{

	return this.reson
}

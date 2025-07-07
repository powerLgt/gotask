package vo

type ResponseVO struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Fail(msg string) (resp ResponseVO) {
	resp = ResponseVO{
		Code: -1,
		Msg:  msg,
		Data: nil,
	}
	return
}

func Success(msg string, data interface{}) (resp ResponseVO) {
	resp = ResponseVO{
		Code: 0,
		Msg:  msg,
		Data: data,
	}
	return
}

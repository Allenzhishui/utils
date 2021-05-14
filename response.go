package utils

// Response  format
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"` //omitempty:如果为空则没有这个字段
}

// page format
// Message
type page struct {
	Count int         `json:"count"`
	Items interface{} `json:"items"`
}

const (
	statusSucc    int = 200 //正常
	statusFail    int = 300 //失败
	statusErrIpt  int = 310 //输入数据有误
	statusErrOpt  int = 320 //无数据返回
	statusErrDeny int = 330 //没有权限
	statusErrJwt  int = 340 //jwt未通过验证
	statusErrSvr  int = 350 //服务端错误
	statusExt     int = 400 //其他约定 //eg 更新 token
)

func newResponse(code int, msg string, data ...interface{}) (int, Response) {
	if len(data) > 0 {
		return 200, Response{
			Code: code,
			Msg:  msg,
			Data: data[0],
		}
	}
	return 200, Response{
		Code: code,
		Msg:  msg,
	}
}

// Succ 返回一个成功标识的结果格式
func Succ(msg string, data ...interface{}) (int, Response) {
	return newResponse(statusSucc, msg, data...)
}

// Fail 返回一个失败标识的结果格式
func Fail(msg string, data ...interface{}) (int, Response) {
	return newResponse(statusFail, msg, data...)
}

// Page 返回一个带有分页数据的结果格式
func Page(msg string, items interface{}, count int) (int, Response) {
	return 200, Response{
		Code: statusSucc,
		Msg:  msg,
		Data: page{
			Items: items,
			Count: count,
		},
	}
}

// ErrIpt 返回一个输入错误的结果格式
func ErrIpt(msg string, data ...interface{}) (int, Response) {
	return newResponse(statusErrIpt, msg, data...)
}

// ErrOpt 返回一个输出错误的结果格式
func ErrOpt(msg string, data ...interface{}) (int, Response) {
	return newResponse(statusErrOpt, msg, data...)
}

// ErrDeny 返回一个没有权限的结果格式
func ErrDeny(msg string, data ...interface{}) (int, Response) {
	return newResponse(statusErrDeny, msg, data...)
}

// ErrJwt 返回一个通过验证的结果格式
func ErrJwt(msg string, data ...interface{}) (int, Response) {
	return newResponse(statusErrJwt, msg, data...)
}

// ErrSvr 返回一个服务端错误的结果格式
func ErrSvr(msg string, data ...interface{}) (int, Response) {
	return newResponse(statusErrSvr, msg, data...)
}

// Ext 返回一个其他约定的结果格式
func Ext(msg string, data ...interface{}) (int, Response) {
	return newResponse(statusExt, msg, data...)
}

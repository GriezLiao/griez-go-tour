package errorcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = make(map[int]string)

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (error *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息:：%s", error.Code(), error.Msg())
}

func (error *Error) Code() int {
	return error.code
}

func (error *Error) Msg() string {
	return error.msg
}

func (error *Error) MsgFormat(args []interface{}) string {
	return fmt.Sprintf(error.msg, args...)
}

func (error *Error) Details() []string {
	return error.details
}

func (error *Error) WithDetails(details ...string) *Error {
	newError := *error
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}

	return &newError
}

func (error *Error) StatusCode() int {
	switch error.code {
	case Success.code:
		return http.StatusOK
	case ServerError.code:
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}

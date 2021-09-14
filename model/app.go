package model

import "fmt"

const (
	DefCode = iota + 100
	ErrCodeBadRequest
	ErrCheckFormData
	ErrServerError
	ErrRecordNotFound
	ErrLogicError
	ErrUserNamePassword
	ErrAuthRequire
	ErrAuthCheckLen
	ErrAuthTokenInvalid
	ErrDbExecError
	ErrDbExec
	ErrNickNameDictWord
	ErrJsonCodeErr
	ErrDataOfflineOrNotFound
)

var msgDefine = map[int]string{
	ErrCodeBadRequest:        "bad request",
	ErrAuthRequire:           "需要附加认证信息",
	ErrAuthCheckLen:          "认证信息格式错误",
	ErrAuthTokenInvalid:      "认证信息失效",
	ErrCheckFormData:         "表单提交错误",
	ErrServerError:           "服务器内部错误",
	ErrDbExecError:           "数据库执行错误",
	ErrRecordNotFound:        "记录不存在",
	ErrUserNamePassword:      "用户名密码错误",
	ErrLogicError:            "业务逻辑错误",
	ErrDbExec:                "数据执行错误",
	ErrNickNameDictWord:      "昵称包含敏感词，请修改",
	ErrDataOfflineOrNotFound: "数据不存在或已下线",
	ErrJsonCodeErr:           "JSON加解密失败",
}

type AppError struct {
	code    int
	message string
}

func (e *AppError) Error() string {
	if e == nil {
		return fmt.Sprintf(" error nil ")
	}
	return fmt.Sprintf("app error. code:%d , Message:%s", e.code, e.message)
}

func (e *AppError) MsgError() string {
	if e == nil {
		return fmt.Sprintf(" error nil ")
	}
	return e.message
}

func NewException(Code int) *AppError {

	message, ok := msgDefine[Code]
	if !ok {
		message = "unknown error"
	}

	e := &AppError{
		code:    Code,
		message: message,
	}

	return e
}

func NewExceptionWithMessage(code int, msg string) *AppError {
	return &AppError{
		code:    code,
		message: msg,
	}
}

/*
 * @Author       : jayj
 * @Date         : 2021-06-26 23:03:32
 * @Description  :
 */
package res

type ErrorCode int

const (
	_ int = iota + 1000
	// OK 成功
	OK
	// ParamsInvalid 参数错误
	ParamsInvalid
	// InternalError 服务器内部错误
	InternalServerError
	// RegisterFailed 注册失败
	RegisterFailed
	// TokenInvalid 无效的token
	TokenInvalid
	// TokenExpired token过期
	TokenExpired
	// UserNotExists 用户不存在
	UserNotExists
	// UrlNotFound 不存在的url
	UrlNotFound
	// FileDontExist 文件不存在
	FileDontExist
	// Unauthorized
	UnauthorizedError
	// InvalidAccountOrPassword
	InvalidAccountOrPassword
	// DuplicatedName
	DuplicatedName
)

var Msg = map[int]string{
	OK:                       "成功",
	ParamsInvalid:            "参数错误",
	InternalServerError:      "服务器内部错误",
	RegisterFailed:           "注册失败",
	TokenInvalid:             "无效的token",
	TokenExpired:             "登录已过期",
	UserNotExists:            "用户不存在",
	UrlNotFound:              "不存在的url",
	FileDontExist:            "文件不存在",
	UnauthorizedError:        "未授权的访问，请先登录或确定有相应的权限",
	InvalidAccountOrPassword: "无效的用户名或密码",
	DuplicatedName:           "用户名已存在",
}

func GetMsg(code int) string {
	return Msg[code]
}

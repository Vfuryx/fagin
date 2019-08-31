package errno

var (
	// 系统错误
	OK                  = &Errno{Code: 0, Message: "OK"}
	NotFound            = &Errno{Code: 10000, Message: "404"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server exception"}
	ErrBind             = &Errno{Code: 10002, Message: "将请求主体绑定到结构时发生错误"}

	// 用户错误
	ErrEncrypt           = &Errno{Code: 20101, Message: "密码加密异常"}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "找不到用户"}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "令牌无效"}
	ErrToken             = &Errno{Code: 20104, Message: "签署JSON Web令牌时出错"}
	ErrPasswordIncorrect = &Errno{Code: 20105, Message: "密码不正确"}
	ErrLogin             = &Errno{Code: 20106, Message: "用户或密码不正确"}
	ErrPermissionDenied  = &Errno{Code: 20107, Message: "权限不足"}

)

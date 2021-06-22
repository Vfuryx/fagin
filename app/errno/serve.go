package errno

type serve struct {
	// 基础错误  10000 ～ 19999
	NotFound,
	SystemERR,
	BindErr,
	RequestTimeoutErr,

	// 上传
	UploadSizeExceededErr,
	UploadFileErr,
	OpenFileErr,
	// 缓存
	CacheIsCloseErr,
	CacheIsNilErr,
	// 增删给查
	ListErr,
	ShowErr,
	StoreErr,
	UpdateErr,
	DeleteErr,

	// 授权错误  20000 ～ 29999
	ErrEncrypt,
	ErrUserNotFound,
	ErrTokenInvalid,
	ErrToken,
	ErrPasswordIncorrect,
	ErrLogin,
	ErrPermissionDenied,
	ErrAuthCheckRole,
	ErrGetCaptcha,
	ErrVerifyCaptcha,
	RoleKeyExistErr,
	RoleRelationExistErr,
	PermissionRelationExistErr,
	MenuRelationExistErr,
	MenuSubExistErr,
	AdminUsernameExistErr,


	//结束
	end *Errno
}

var Serve = serve{
	// 系统错误
	NotFound:          &Errno{Code: 10000, Message: "404"},
	SystemERR:         &Errno{Code: 10001, Message: "系统错误"},
	BindErr:           &Errno{Code: 10002, Message: "参数错误"},
	RequestTimeoutErr: &Errno{Code: 10003, Message: "请求超时"},

	// 上传
	UploadSizeExceededErr: &Errno{Code: 10100, Message: "超出上传文件预定最大值"},
	UploadFileErr:         &Errno{Code: 10101, Message: "上传文件错误"},
	OpenFileErr:           &Errno{Code: 10102, Message: "打开文件失败"},
	// 缓存
	CacheIsCloseErr: &Errno{Code: 10200, Message: "缓存服务未开启"},
	CacheIsNilErr:   &Errno{Code: 10201, Message: "缓存不存在"},
	// 增删给查
	ListErr:   &Errno{Code: 10301, Message: "获取数据失败"},
	ShowErr:   &Errno{Code: 10302, Message: "获取详情失败"},
	StoreErr:  &Errno{Code: 10303, Message: "新增数据失败"},
	UpdateErr: &Errno{Code: 10304, Message: "更新数据失败"},
	DeleteErr: &Errno{Code: 10305, Message: "删除数据失败"},


	// 用户错误
	ErrEncrypt:           &Errno{Code: 20001, Message: "密码加密异常"},
	ErrUserNotFound:      &Errno{Code: 20002, Message: "找不到用户"},
	ErrTokenInvalid:      &Errno{Code: 20003, Message: "令牌无效"},
	ErrToken:             &Errno{Code: 20004, Message: "签署JSON Web令牌时出错"},
	ErrPasswordIncorrect: &Errno{Code: 20005, Message: "密码不正确"},
	ErrLogin:             &Errno{Code: 20006, Message: "用户或密码不正确"},
	ErrPermissionDenied:  &Errno{Code: 20007, Message: "权限不足"},

	ErrAuthCheckRole: &Errno{Code: 20012, Message: "对不起，您没有该接口访问权限，请联系管理员"},
	ErrGetCaptcha:    &Errno{Code: 20013, Message: "获取验证码失败"},
	ErrVerifyCaptcha: &Errno{Code: 20014, Message: "验证码验证失败"},

	RoleKeyExistErr:            &Errno{Code: 21001, Message: "角色关键字已存在"},
	RoleRelationExistErr:       &Errno{Code: 21002, Message: "角色存在关联"},
	PermissionRelationExistErr: &Errno{Code: 21003, Message: "权限存在关联"},
	MenuRelationExistErr:       &Errno{Code: 21004, Message: "菜单存在关联"},
	MenuSubExistErr:            &Errno{Code: 21005, Message: "菜单存在下级"},
	AdminUsernameExistErr:      &Errno{Code: 21006, Message: "用户名已存在"},
}

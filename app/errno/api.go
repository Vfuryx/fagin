package errno

var (
	// 系统错误
	OK                    = &Errno{Code: 0, Message: "OK"}
	NotFound              = &Errno{Code: 10000, Message: "404"}
	InternalServerError   = &Errno{Code: 10001, Message: "Internal server exception"}
	ErrBind               = &Errno{Code: 10002, Message: "将请求主体绑定到结构时发生错误/参数错误"}
	ErrUploadSizeExceeded = &Errno{Code: 10003, Message: "超出上传文件预定最大值"}
	ErrUploadFile         = &Errno{Code: 10004, Message: "上传文件错误"}
	ErrOpenFile           = &Errno{Code: 10005, Message: "打开文件失败"}
	ErrCacheIsClose       = &Errno{Code: 10006, Message: "缓存服务未开启"}
	ErrCacheIsNil         = &Errno{Code: 10006, Message: "缓存不存在"}

	// 用户错误
	ErrEncrypt           = &Errno{Code: 20101, Message: "密码加密异常"}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "找不到用户"}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "令牌无效"}
	ErrToken             = &Errno{Code: 20104, Message: "签署JSON Web令牌时出错"}
	ErrPasswordIncorrect = &Errno{Code: 20105, Message: "密码不正确"}
	ErrLogin             = &Errno{Code: 20106, Message: "用户或密码不正确"}
	ErrPermissionDenied  = &Errno{Code: 20107, Message: "权限不足"}
	ErrAddUser           = &Errno{Code: 20108, Message: "创建用户失败"}
	ErrUpdateUser        = &Errno{Code: 20109, Message: "更新用户失败"}
	ErrDeleteUser        = &Errno{Code: 20109, Message: "删除用户失败"}
	ErrUserList          = &Errno{Code: 20110, Message: "获取用户列表失败"}

	// video err
	ErrAddVideo    = &Errno{Code: 20201, Message: "创建视频失败"}
	ErrUpdateVideo = &Errno{Code: 20202, Message: "更新视频失败"}
	ErrDeleteVideo = &Errno{Code: 20203, Message: "删除视频失败"}
	ErrVideoList   = &Errno{Code: 20204, Message: "获取视频列表失败"}

	// banner err
	ErrAddBanner    = &Errno{Code: 20211, Message: "创建轮播图失败"}
	ErrUpdateBanner = &Errno{Code: 20212, Message: "更新轮播图失败"}
	ErrDeleteBanner = &Errno{Code: 20213, Message: "删除轮播图失败"}
	ErrBannerList   = &Errno{Code: 20214, Message: "获取轮播图列表失败"}
	ErrBanner       = &Errno{Code: 20215, Message: "获取轮播图失败"}

	// 网站设置 err
	ErrUpdateWebsiteConfig = &Errno{Code: 20221, Message: "更新网站配置失败"}
	ErrWebsiteConfig       = &Errno{Code: 20221, Message: "获取网站配置失败"}

	// 公司介绍 err
	ErrUpdateCompanyIntroduction = &Errno{Code: 20231, Message: "更新公司介绍失败"}
	ErrCompanyIntroduction       = &Errno{Code: 20231, Message: "获取公司介绍失败"}
)

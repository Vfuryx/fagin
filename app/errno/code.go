package errno

// Errno 错误码
type Errno int

// Get 获取错误码
func (e Errno) Get() int {
	return int(e)
}

const (
	// OK 基础错误  0 ～ 999
	OK                Errno = 0
	InternalServerErr Errno = 500
	NotFoundErr       Errno = 404

	// SysErr 1000 ～ 9999 系统 错误
	SysErr Errno = 1000

	// 第一位是 错误位置 1:请求验证; 2:中间件 3:控制器; 4:服务; 5:DAO; 6:模型; 7:缓存; 8:定时任务 9:队列;
	// 第二到三位是 错误模块 01:用户; 02:产品; 03:订单 09:文件;

	// ReqErr 请求验证错误  100000 ～ 199999
	ReqErr                   Errno = 100000
	ReqUploadFileErr         Errno = 100001
	ReqUploadSizeExceededErr Errno = 100002
	ReqRequestTimeoutErr     Errno = 100003
	ReqMenuNotFoundErr       Errno = 101015
	ReqStructureNotFoundErr  Errno = 101016

	// MidErr  中间件错误  100000 ～ 199999
	MidErr                 Errno = 200000
	MidAuthCheckRoleErr    Errno = 201001
	MidPermissionDeniedErr Errno = 201002

	// CtxErr  控制器错误 300000 ～ 399999
	CtxErr              Errno = 300000
	CtxListErr          Errno = 300001
	CtxShowErr          Errno = 300002
	CtxStoreErr         Errno = 300003
	CtxUpdateErr        Errno = 300004
	CtxDeleteErr        Errno = 300005
	CtxUserNotFoundErr  Errno = 301001
	CtxGetCaptchaErr    Errno = 301002
	CtxVerifyCaptchaErr Errno = 301003
	CtxTokenInvalidErr  Errno = 301004
	CtxRoleKeyExistErr  Errno = 301005
	CtxUserExistErr     Errno = 301006
	CtxOpenFileErr      Errno = 309001

	// SerErr 服务错误 400000 ～ 499999
	SerErr                        Errno = 400000
	SerTokenErr                   Errno = 401001
	SerTokenNotFoundErr           Errno = 401002
	SerTokenInvalidErr            Errno = 401003
	SerTokenExpiredErr            Errno = 401004
	SerTokenGuardErr              Errno = 401005
	SerAccountOrPasswordErr       Errno = 401006
	SerMenuSubExistErr            Errno = 401007
	SerMenuRelationExistErr       Errno = 401008
	SerRoleRelationExistErr       Errno = 401009
	SerPermissionRelationExistErr Errno = 401010
	SerMenuPathsUnsafeErr         Errno = 401011

	// DaoErr 500000 ～ 599999
	DaoErr                Errno = 500000
	DaoMenuPathsUnsafeErr Errno = 501001

	// ModErr 600000 ～ 699999
	ModErr Errno = 600000

	// CacheErr 700000 ～ 799999
	CacheErr        Errno = 700000
	CacheIsCloseErr Errno = 700001
	CacheIsNilErr   Errno = 700002
)

const defErr = "未知错误"

var errorMap = map[Errno]string{
	OK:                "OK",
	InternalServerErr: "Internal server exception",
	NotFoundErr:       "404",

	SysErr: "系统错误",

	ReqErr:                   "请求参数错误",
	ReqUploadFileErr:         "上传文件错误",
	ReqUploadSizeExceededErr: "超出上传文件预定最大值",
	ReqRequestTimeoutErr:     "请求超时",
	ReqMenuNotFoundErr:       "菜单不存在",
	ReqStructureNotFoundErr:  "构造不存在",

	MidErr:                 "中间件错误",
	MidAuthCheckRoleErr:    "对不起，您没有该接口访问权限，请联系管理员",
	MidPermissionDeniedErr: "无权限",

	CtxErr:              "控制器错误",
	CtxListErr:          "获取数据失败",
	CtxShowErr:          "获取详情失败",
	CtxStoreErr:         "新增数据失败",
	CtxUpdateErr:        "更新数据失败",
	CtxDeleteErr:        "删除数据失败",
	CtxUserNotFoundErr:  "用户不存在",
	CtxUserExistErr:     "用户已存在",
	CtxGetCaptchaErr:    "获取验证码失败",
	CtxVerifyCaptchaErr: "验证码错误",
	CtxTokenInvalidErr:  "授权失败",
	CtxRoleKeyExistErr:  "角色关键字已存在",
	CtxOpenFileErr:      "打开文件失败",

	SerErr:                        "服务错误",
	SerTokenErr:                   "令牌错误",
	SerTokenNotFoundErr:           "令牌不存在",
	SerTokenInvalidErr:            "令牌无效",
	SerTokenExpiredErr:            "令牌过期",
	SerTokenGuardErr:              "令牌跨界",
	SerAccountOrPasswordErr:       "账号或密码错误",
	SerMenuRelationExistErr:       "菜单存在关联",
	SerMenuSubExistErr:            "菜单存在下级",
	SerRoleRelationExistErr:       "角色存在关联",
	SerPermissionRelationExistErr: "权限存在关联",
	SerMenuPathsUnsafeErr:         "菜单层级路径非法定义",

	DaoErr:                "DAO错误",
	DaoMenuPathsUnsafeErr: "菜单层级路径非法定义",

	ModErr: "模型错误",

	CacheErr:        "缓存错误",
	CacheIsCloseErr: "缓存服务未开启",
	CacheIsNilErr:   "缓存不存在",
}

func (e Errno) Error() string {
	if msg, ok := errorMap[e]; ok {
		return msg
	}

	return defErr
}

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
	SerPasswordIncorrectErr       Errno = 401006
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

func (e Errno) Error() string {
	switch e {
	case OK:
		return "OK"
	case InternalServerErr:
		return "Internal server exception"
	case NotFoundErr:
		return "404"

	case SysErr:
		return "系统错误"

	case ReqErr:
		return "请求参数错误"
	case ReqUploadFileErr:
		return "上传文件错误"
	case ReqUploadSizeExceededErr:
		return "超出上传文件预定最大值"
	case ReqRequestTimeoutErr:
		return "请求超时"
	case ReqMenuNotFoundErr:
		return "菜单不存在"
	case ReqStructureNotFoundErr:
		return "构造不存在"

	case MidErr:
		return "中间件错误"
	case MidAuthCheckRoleErr:
		return "对不起，您没有该接口访问权限，请联系管理员"
	case MidPermissionDeniedErr:
		return "无权限"

	case CtxErr:
		return "控制器错误"
	case CtxListErr:
		return "获取数据失败"
	case CtxShowErr:
		return "获取详情失败"
	case CtxStoreErr:
		return "新增数据失败"
	case CtxUpdateErr:
		return "更新数据失败"
	case CtxDeleteErr:
		return "删除数据失败"
	case CtxUserNotFoundErr:
		return "找不到用户"
	case CtxUserExistErr:
		return "用户已存在"
	case CtxGetCaptchaErr:
		return "获取验证码失败"
	case CtxVerifyCaptchaErr:
		return "验证码错误"
	case CtxTokenInvalidErr:
		return "授权失败"
	case CtxRoleKeyExistErr:
		return "角色关键字已存在"
	case CtxOpenFileErr:
		return "打开文件失败"

	case SerErr:
		return "服务错误"
	case SerTokenErr:
		return "令牌错误"
	case SerTokenNotFoundErr:
		return "令牌不存在"
	case SerTokenInvalidErr:
		return "令牌无效"
	case SerTokenExpiredErr:
		return "令牌过期"
	case SerTokenGuardErr:
		return "令牌跨界"
	case SerPasswordIncorrectErr:
		return "密码不正确"
	case SerMenuRelationExistErr:
		return "菜单存在关联"
	case SerMenuSubExistErr:
		return "菜单存在下级"
	case SerRoleRelationExistErr:
		return "角色存在关联"
	case SerPermissionRelationExistErr:
		return "权限存在关联"
	case SerMenuPathsUnsafeErr:
		return "菜单层级路径非法定义"

	case DaoErr:
		return "DAO错误"
	case DaoMenuPathsUnsafeErr:
		return "菜单层级路径非法定义"

	case ModErr:
		return "模型错误"

	case CacheErr:
		return "缓存错误"
	case CacheIsCloseErr:
		return "缓存服务未开启"
	case CacheIsNilErr:
		return "缓存不存在"

	default:
		return "未知错误"
	}
}

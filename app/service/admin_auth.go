package service

import (
	"fagin/app"
	"fagin/app/enums"
	"fagin/app/errno"
	"fagin/app/models/admin_user"
	"fagin/config"
	"fagin/pkg/errorw"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

type adminAuthService struct{}

// AdminAuth 后台授权服务
var AdminAuth adminAuthService

// Login 登录
func (*adminAuthService) Login(name string, password string) (uint, error) {
	params := map[string]interface{}{
		"username": name,
		"status":   enums.StatusActive.Get(),
	}
	columns := []string{"id", "username", "password", "status"}
	adminUser := admin_user.New()
	if err := adminUser.Dao().Query(params, columns, nil).First(&adminUser); err != nil {
		return 0, errno.CtxUserNotFoundErr
	}

	if adminUser.Username != name {
		return 0, errno.SerPasswordIncorrectErr
	}

	//匹配密码
	if err := app.Compare(adminUser.Password, password); err != nil {
		return 0, errno.SerPasswordIncorrectErr
	}

	adminUser.Dao().Login(adminUser.ID)

	return adminUser.ID, nil
}

// Logout 登出
func (*adminAuthService) Logout(id uint) error {
	err := admin_user.NewDao().SetCheckInAt(id, time.Now().Unix())
	return errorw.UP(err)
}

// AdminCustomClaims 自定义Claims
type AdminCustomClaims struct {
	UserID uint   `json:"user_id,omitempty"`
	Guard  string `json:"guard,omitempty"`
	*jwt.StandardClaims
}

// Valid 自定义验证方法
func (cc AdminCustomClaims) Valid() error {
	if cc.Guard != admin_user.AdminUserGuard {
		return errno.SerTokenExpiredErr
	}
	// StandardClaims 基础验证方法
	return cc.StandardClaims.Valid()
}

// Sign signs the context with the specified secret.
func (*adminAuthService) Sign(UserID uint, AdminUser, secret string) (tokenString string, err error) {
	// Load the jwt secret from the Gin config if the secret isn't specified.
	if secret == "" {
		secret = config.App().Key
	}
	now := time.Now()
	// 期限
	maxAge := 60 * 60 * 24
	// 自定义Claims
	customClaims := &AdminCustomClaims{
		UserID: UserID,                    // 用户id
		Guard:  admin_user.AdminUserGuard, // Guard
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: now.Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间，必须设置
			Issuer:    AdminUser,                                           // 非必须，也可以填充用户名，
			IssuedAt:  now.Unix(),                                          // 必须，签发时间，
		},
	}

	// The token content.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))
	return tokenString, errorw.UP(err)
}

//// RefreshToken 更新 token
//func (auth *adminAuthService) RefreshToken(token string) (string, error) {
//	jwt.TimeFunc = func() time.Time {
//		return time.Unix(0, 0)
//	}
//	t, err := jwt.ParseWithClaims(token, &AdminCustomClaims{}, auth.secretFunc(config.App().Key))
//	if err != nil {
//		return "", errno.SerTokenInvalidErr
//	}
//	if claims, ok := t.Claims.(*AdminCustomClaims); ok && t.Valid {
//		jwt.TimeFunc = time.Now
//		return auth.Sign(claims.UserID, claims.Issuer, config.App().Key)
//	}
//	return "", errno.SerTokenErr
//}

type adminContext struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}

// AdminLogin 后台登录上下文
var AdminLogin = &adminContext{}

func (auth *adminAuthService) ParseRequest(ctx *gin.Context) (*adminContext, error) {
	head := ctx.Request.Header.Get("Authorization")
	prefix := "Bearer "

	// 空
	if head == "" {
		return AdminLogin, errno.SerTokenNotFoundErr
	}
	// 格式不对
	if !strings.HasPrefix(head, prefix) {
		return AdminLogin, errno.SerTokenErr
	}
	// 截取token
	token := head[strings.Index(head, prefix)+len(prefix):]

	return auth.Parse(token, config.App().Key)
}

func (auth *adminAuthService) Parse(token string, secret string) (ctx *adminContext, err error) {
	ctx = AdminLogin
	t, err := jwt.ParseWithClaims(token, &AdminCustomClaims{}, auth.secretFunc(secret))
	if err != nil {
		return ctx, errno.SerTokenInvalidErr
	}
	if claims, ok := t.Claims.(*AdminCustomClaims); ok && t.Valid {
		// 验证签发时间
		adminUser := admin_user.New()
		err = adminUser.Dao().FindById(uint(claims.UserID), []string{"id", "check_in_at"})
		if err != nil {
			return ctx, err
		}
		// token 签发时间 小于 用户最后签发的时间 代表 token 过期
		if claims.IssuedAt < int64(adminUser.CheckInAt) {
			return ctx, errno.SerTokenExpiredErr
		}
		ctx.Name = claims.Issuer
		ctx.UserID = claims.UserID
		return ctx, nil
	}
	return ctx, errno.SerTokenErr
}

// secretFunc validates the secret format.
func (*adminAuthService) secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}

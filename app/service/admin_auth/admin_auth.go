package admin_auth

import (
	"fagin/app/enums"
	"fagin/app/errno"
	"fagin/app/models/admin_user"
	"fagin/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type adminAuthService struct{}

var AdminAuth adminAuthService

// Login 登录
func (*adminAuthService) Login(name string, password string) (uint64, error) {
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
	if err := Compare(adminUser.Password, password); err != nil {
		return 0, errno.SerPasswordIncorrectErr
	}

	adminUser.Dao().Login(adminUser.ID)

	return uint64(adminUser.ID), nil
}

// Logout 登出
func (*adminAuthService) Logout(id uint) error {
	return admin_user.NewDao().SetCheckInAt(id, time.Now().Unix())
}

// Encrypt encrypts the plain text with bcrypt.
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// Compare compares the encrypted text with the plain text if it's the same.
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// CustomClaims 自定义Claims
type CustomClaims struct {
	UserID uint64 `json:"user_id,omitempty"`
	Guard  string `json:"guard,omitempty"`
	*jwt.StandardClaims
}

// Valid 自定义验证方法
func (cc CustomClaims) Valid() error {
	if cc.Guard != admin_user.AdminUserGuard {
		return errno.SerTokenExpiredErr
	}
	// StandardClaims 基础验证方法
	return cc.StandardClaims.Valid()
}

// Sign signs the context with the specified secret.
func (*adminAuthService) Sign(UserID uint64, AdminUser, secret string) (tokenString string, err error) {
	// Load the jwt secret from the Gin config if the secret isn't specified.
	if secret == "" {
		secret = config.App().Key
	}
	now := time.Now()
	// 期限
	maxAge := 60 * 60 * 24
	// 自定义Claims
	customClaims := &CustomClaims{
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
	return
}

type context struct {
	UserID uint64 `json:"user_id"`
	Name   string `json:"name"`
}

var Login = &context{}

func (auth *adminAuthService) ParseRequest(ctx *gin.Context) (*context, error) {
	head := ctx.Request.Header.Get("Authorization")
	prefix := "Bearer "

	// 空
	if head == "" {
		return Login, errno.SerTokenNotFoundErr
	}
	// 格式不对
	if !strings.HasPrefix(head, prefix) {
		return Login, errno.SerTokenErr
	}
	// 截取token
	token := head[strings.Index(head, prefix)+len(prefix):]

	return auth.Parse(token, config.App().Key)
}

func (auth *adminAuthService) Parse(token string, secret string) (ctx *context, err error) {
	ctx = Login
	t, err := jwt.ParseWithClaims(token, &CustomClaims{}, auth.secretFunc(secret))
	if err != nil {
		return ctx, errno.SerTokenInvalidErr
	}
	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
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
	} else {
		return ctx, errno.SerTokenErr
	}
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

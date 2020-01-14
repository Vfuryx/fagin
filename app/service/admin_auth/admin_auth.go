package admin_auth

import (
	"fagin/app/constants/status"
	"fagin/app/errno"
	"fagin/app/models/admin_user"
	"fagin/config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type adminAuthService struct{}

var AdminAuth adminAuthService

/**
 * 登录
 */
func (adminAuthService) Login(name string, password string) error {
	params := map[string]interface{}{
		"username": name,
		"status":   status.Active,
	}
	columns := []string{"id", "username", "password", "status"}
	adminUser := admin_user.New()
	if err := adminUser.Dao().Query(params, columns, nil).First(&adminUser); err != nil {
		return errno.InternalServerError
	}

	if adminUser.Username != name {
		return errno.ErrPasswordIncorrect
	}

	//匹配密码
	if err := Compare(adminUser.Password, password); err != nil {
		return errno.ErrPasswordIncorrect
	}

	return nil
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

// Sign signs the context with the specified secret.
func (adminAuthService) Sign(AdminUser, secret string) (tokenString string, err error) {
	// Load the jwt secret from the Gin config if the secret isn't specified.
	if secret == "" {
		secret = config.App.Key
	}
	// The token content.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": AdminUser,
		"nbf":  time.Now().Unix(),
		"iat":  time.Now().Unix(),
	})
	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	return
}

type context struct {
	Name string `json:"name"`
}

var Login = &context{}

func (auth adminAuthService) ParseRequest(ctx *gin.Context) (*context, error) {
	head := ctx.Request.Header.Get("Authorization")

	if len(head) == 0 {
		return Login, errno.ErrTokenInvalid
	}

	var token string
	fmt.Sscanf(head, "Bearer %s", &token)

	return auth.Parse(token, config.App.Key)
}

func (auth adminAuthService) Parse(token string, secret string) (*context, error) {
	ctx := Login
	if token, err := jwt.Parse(token, auth.secretFunc(secret)); err != nil {
		return ctx, errno.ErrTokenInvalid
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Name = claims["name"].(string)
		return ctx, nil
	} else {
		return ctx, errno.ErrTokenInvalid
	}
}

// secretFunc validates the secret format.
func (adminAuthService) secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

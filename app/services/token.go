package services

import (
	"fagin/app/errno"
	"fagin/app/models/user"
	"fagin/config"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type tokenService struct{}

type context struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

var Token tokenService
var Login = &context{}

func (ts *tokenService) ParseRequest(ctx *gin.Context) (*context, error) {
	head := ctx.Request.Header.Get("Authorization")

	if head == "" {
		return Login, errno.SerTokenNotFoundErr
	}

	var token string

	_, err := fmt.Sscanf(head, "Bearer %s", &token)
	if err != nil {
		return nil, err
	}

	return ts.Parse(token, config.JWT().Secret())
}

func (ts *tokenService) Parse(token, secret string) (*context, error) {
	ctx := Login
	if token, err := jwt.Parse(token, ts.secretFunc(secret)); err != nil {
		return ctx, errno.SerTokenInvalidErr
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = uint64(claims["id"].(float64))
		if ctx.Username, ok = claims["username"].(string); !ok {
			return nil, errno.SerTokenInvalidErr
		}

		return ctx, nil
	} else {
		return ctx, errno.SerTokenInvalidErr
	}
}

// secretFunc validates the secret format.
func (ts *tokenService) secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// Sign signs the context with the specified secret.
func (ts *tokenService) Sign(u *user.User, secret string) (tokenString string, err error) {
	// Load the jwt secret from the Gin config if the secret isn't specified.
	if secret == "" {
		secret = config.JWT().Secret()
	}
	// The token content.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       u.ID,
		"username": u.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})
	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	return
}

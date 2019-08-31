package service

import (
	"fagin/app/errno"
	"fagin/app/models/user"
	"fagin/config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type token struct{}

type context struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

var Token = token{}
var Login = &context{}

func (this token) ParseRequest(ctx *gin.Context) (*context, error) {
	head := ctx.Request.Header.Get("Authorization")

	secret := config.Jwt.Secret

	if len(head) == 0 {
		return Login, errno.ErrTokenInvalid
	}

	var token string
	fmt.Sscanf(head, "Bearer %s", &token)

	return this.Parse(token, secret)
}

func (this token) Parse(token string, secret string) (*context, error) {
	ctx := Login
	if token, err := jwt.Parse(token, this.secretFunc(secret)); err != nil {
		return ctx, errno.ErrTokenInvalid
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = uint64(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		return ctx, nil
	} else {
		return ctx, errno.ErrTokenInvalid
	}
}

// secretFunc validates the secret format.
func (this token) secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// Sign signs the context with the specified secret.
func (this token) Sign(user user.User, secret string) (tokenString string, err error) {
	// Load the jwt secret from the Gin config if the secret isn't specified.
	if secret == "" {
		secret = config.Jwt.Secret
	}
	// The token content.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})
	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	return
}

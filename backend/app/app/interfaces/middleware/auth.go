package middleware

import (
	"context"
	"fmt"
	"github.com/lestrrat-go/jwx/jwt"
	"net/http"
	"todo-list/app/config"
)

type contextKey string

const userCtxKey contextKey = "user"

type User struct {
	AuthID string
	Name   string
	Email  string
}

type AuthMiddleware interface {
	Handler(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request, error)
}

type authMiddleware struct {
	conf *config.Config
}

func NewAuthMiddleware(conf *config.Config) (AuthMiddleware, error) {
	return &authMiddleware{conf: conf}, nil
}

func (m *authMiddleware) Handler(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request, error) {
	ctx := r.Context()

	// IDトークンの検証
	idToken := r.Header.Get("Authorization")
	token, err := jwt.ParseString(idToken)
	if err != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return w, r, err
	}

	// JWTの検証
	issuer := fmt.Sprintf(
		"https://cognito-idp.%s.amazonaws.com/%s",
		*m.conf.AWS.Config.Region,
		m.conf.AWS.UserPoolID,
	)
	err = jwt.Validate(token, jwt.WithAudience(m.conf.AWS.UserPoolClientID), jwt.WithIssuer(issuer))
	if err != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return w, r, err
	}

	// TODO ユーザの存在チェック(Signinの場合はチェックしない)

	name := ""
	email := ""
	if value, ok := token.Get("cognito:username"); ok {
		name = value.(string)
	}
	if value, ok := token.Get("email"); ok {
		email = value.(string)
	}
	user := &User{
		AuthID: token.Subject(),
		Name:   name,
		Email:  email,
	}
	ctx = SetUser(ctx, user)

	r = r.WithContext(ctx)
	w.Header().Set("Content-Type", "application/json")
	return w, r, nil
}

func SetUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userCtxKey, user)
}

func UserForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(userCtxKey).(*User)
	return raw
}

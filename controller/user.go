package controller

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"

	"goa.design/goa/v3/security"

	"goa-shopping/config"
	"goa-shopping/dao"
	auth "goa-shopping/gen/auth"
	"goa-shopping/gen/log"
)

var RequestIDKey int

// auth service example implementation.
// The example methods log the requests and return zero values.
type authsrvc struct {
	logger *log.Logger
}

// NewAuth returns the auth service implementation.
func NewAuth(logger *log.Logger) auth.Service {
	return &authsrvc{logger}
}

// JWTAuth implements the authorization logic for service "auth" for the "jwt"
// security scheme.
func (s *authsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	// logger := L(ctx, a.logger)

	claims := make(jwt.MapClaims)

	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	if _, err := jwt.ParseWithClaims(token, claims, func(_ *jwt.Token) (interface{}, error) {
		return []byte(config.C.Jwt.Secret), nil
	}); err != nil {
		return ctx, MakeUnauthorizedError(ctx, "invalid token")
	}

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		return ctx, MakeUnauthorizedError(ctx, "invalid scopes in token")
	}

	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return ctx, MakeUnauthorizedError(ctx, "无操作权限")
	}

	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}

	if err := s.validateScopes(scheme.RequiredScopes, scopesInToken); err != nil {
		return ctx, MakeForbiddenError(ctx, err.Error())
	}

	currentUserID, ok := claims["jti"].(string)
	if !ok {
		return ctx, MakeUnauthorizedError(ctx, "无操作权限")
	}

	// 保存用户ID
	ctx = context.WithValue(ctx, RequestIDKey, currentUserID)

	return ctx, nil
}

// create JWT token
func (s *authsrvc) createJwtToken(userID string, userType int, scopes []string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jti":    userID,
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Second * time.Duration(config.C.Jwt.ExpireIn)).Unix(),
		"type":   userType,
		"scopes": scopes,
	})

	// note that if "SignedString" returns an error then it is returned as
	// an internal error to the client
	return token.SignedString([]byte(config.C.Jwt.Secret))
}

func (s *authsrvc) GetCurrentUserID(ctx context.Context) (string, error) {
	if ctx != nil {
		tifUid, _ := ctx.Value(RequestIDKey).(string)
		return tifUid, nil
	}
	return "", errors.New("ctx is nil")
}

func (s *authsrvc) validateScopes(expected, actual []string) error {
	for _, r := range expected {
		found := false
		for _, s := range actual {
			if s == r {
				found = true
				break
			}
		}
		if found {
			return nil
		}
	}
	return fmt.Errorf("您没有权限进行此操作")
}

// 使用手机号或者邮箱 + 密码登录
func (s *authsrvc) Login(ctx context.Context, p *auth.LoginPayload) (res *auth.LoginResult, err error) {
	res = &auth.LoginResult{}
	s.logger.Info("auth.Login")
	return
}

// 获取用户信息
func (s *authsrvc) GetUserInfo(ctx context.Context, p *auth.GetUserInfoPayload) (res *auth.UserInfo, err error) {
	res = &auth.UserInfo{}
	s.logger.Info("auth.GetUserInfo")
	return
}

// 用户注册使用手机/邮箱+验证码注册
func (s *authsrvc) SignUp(ctx context.Context, p *auth.SignUpPayload) (res *auth.Session, err error) {
	res = &auth.Session{}
	s.logger.Info("auth.SignUp")

	userDao := UserDaoFunc(ctx, dao.DB, s.logger)
	_, err = userDao.CreateUser(p)
	return
}

// 重置密码
func (s *authsrvc) SetPassword(ctx context.Context, p *auth.SetPasswordPayload) (res *auth.Success, err error) {
	res = &auth.Success{}
	s.logger.Info("auth.SetPassword")
	return
}

// 请求电子邮箱验证码
func (s *authsrvc) RequestEmailCode(ctx context.Context, p *auth.RequestEmailCodePayload) (res *auth.Success, err error) {
	res = &auth.Success{}
	s.logger.Info("auth.RequestEmailCode")
	return
}

// 请求短信验证码
func (s *authsrvc) RequestSmSCode(ctx context.Context, p *auth.RequestSmSCodePayload) (res *auth.Success, err error) {
	res = &auth.Success{}
	s.logger.Info("auth.RequestSmSCode")
	return
}

var UserDaoFunc = func(ctx context.Context, tx *gorm.DB, logger *log.Logger) dao.UserDaoImpl {
	return dao.NewUserDao(ctx, tx, logger)
}

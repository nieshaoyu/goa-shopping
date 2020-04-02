// Code generated by goa v3.1.1, DO NOT EDIT.
//
// auth endpoints
//
// Command:
// $ goa gen goa-shopping/design

package auth

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "auth" service endpoints.
type Endpoints struct {
	Login            goa.Endpoint
	GetUserInfo      goa.Endpoint
	SignUp           goa.Endpoint
	SetPassword      goa.Endpoint
	RequestEmailCode goa.Endpoint
	RequestSmSCode   goa.Endpoint
}

// NewEndpoints wraps the methods of the "auth" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Login:            NewLoginEndpoint(s),
		GetUserInfo:      NewGetUserInfoEndpoint(s, a.JWTAuth),
		SignUp:           NewSignUpEndpoint(s),
		SetPassword:      NewSetPasswordEndpoint(s),
		RequestEmailCode: NewRequestEmailCodeEndpoint(s),
		RequestSmSCode:   NewRequestSmSCodeEndpoint(s),
	}
}

// Use applies the given middleware to all the "auth" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Login = m(e.Login)
	e.GetUserInfo = m(e.GetUserInfo)
	e.SignUp = m(e.SignUp)
	e.SetPassword = m(e.SetPassword)
	e.RequestEmailCode = m(e.RequestEmailCode)
	e.RequestSmSCode = m(e.RequestSmSCode)
}

// NewLoginEndpoint returns an endpoint function that calls the method "Login"
// of service "auth".
func NewLoginEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*LoginPayload)
		return s.Login(ctx, p)
	}
}

// NewGetUserInfoEndpoint returns an endpoint function that calls the method
// "GetUserInfo" of service "auth".
func NewGetUserInfoEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetUserInfoPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"role:user", "role:admin"},
			RequiredScopes: []string{"role:user"},
		}
		var token string
		if p.JWTToken != nil {
			token = *p.JWTToken
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.GetUserInfo(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedUserInfo(res, "default")
		return vres, nil
	}
}

// NewSignUpEndpoint returns an endpoint function that calls the method
// "SignUp" of service "auth".
func NewSignUpEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SignUpPayload)
		res, err := s.SignUp(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSession(res, "default")
		return vres, nil
	}
}

// NewSetPasswordEndpoint returns an endpoint function that calls the method
// "SetPassword" of service "auth".
func NewSetPasswordEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SetPasswordPayload)
		res, err := s.SetPassword(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSuccess(res, "default")
		return vres, nil
	}
}

// NewRequestEmailCodeEndpoint returns an endpoint function that calls the
// method "RequestEmailCode" of service "auth".
func NewRequestEmailCodeEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RequestEmailCodePayload)
		res, err := s.RequestEmailCode(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSuccess(res, "default")
		return vres, nil
	}
}

// NewRequestSmSCodeEndpoint returns an endpoint function that calls the method
// "RequestSmSCode" of service "auth".
func NewRequestSmSCodeEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RequestSmSCodePayload)
		res, err := s.RequestSmSCode(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSuccess(res, "default")
		return vres, nil
	}
}

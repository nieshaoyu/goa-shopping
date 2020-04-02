// Code generated by goa v3.1.1, DO NOT EDIT.
//
// auth HTTP server encoders and decoders
//
// Command:
// $ goa gen goa-shopping/design

package server

import (
	"context"
	auth "goa-shopping/gen/auth"
	authviews "goa-shopping/gen/auth/views"
	"io"
	"net/http"
	"strings"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeLoginResponse returns an encoder for responses returned by the auth
// Login endpoint.
func EncodeLoginResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*auth.LoginResult)
		enc := encoder(ctx, w)
		body := NewLoginResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeLoginRequest returns a decoder for requests sent to the auth Login
// endpoint.
func DecodeLoginRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body LoginRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateLoginRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewLoginPayload(&body)

		return payload, nil
	}
}

// EncodeGetUserInfoResponse returns an encoder for responses returned by the
// auth GetUserInfo endpoint.
func EncodeGetUserInfoResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*authviews.UserInfo)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
		enc := encoder(ctx, w)
		body := NewGetUserInfoResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetUserInfoRequest returns a decoder for requests sent to the auth
// GetUserInfo endpoint.
func DecodeGetUserInfoRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			jwtToken *string
		)
		jwtTokenRaw := r.Header.Get("Authorization")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		payload := NewGetUserInfoPayload(jwtToken)
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeSignUpResponse returns an encoder for responses returned by the auth
// SignUp endpoint.
func EncodeSignUpResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*authviews.Session)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
		enc := encoder(ctx, w)
		body := NewSignUpResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeSignUpRequest returns a decoder for requests sent to the auth SignUp
// endpoint.
func DecodeSignUpRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body SignUpRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateSignUpRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewSignUpPayload(&body)

		return payload, nil
	}
}

// EncodeSignUpError returns an encoder for errors returned by the SignUp auth
// endpoint.
func EncodeSignUpError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewSignUpBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeSetPasswordResponse returns an encoder for responses returned by the
// auth SetPassword endpoint.
func EncodeSetPasswordResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*authviews.Success)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
		enc := encoder(ctx, w)
		body := NewSetPasswordResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeSetPasswordRequest returns a decoder for requests sent to the auth
// SetPassword endpoint.
func DecodeSetPasswordRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body SetPasswordRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateSetPasswordRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewSetPasswordPayload(&body)

		return payload, nil
	}
}

// EncodeRequestEmailCodeResponse returns an encoder for responses returned by
// the auth RequestEmailCode endpoint.
func EncodeRequestEmailCodeResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*authviews.Success)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
		enc := encoder(ctx, w)
		body := NewRequestEmailCodeResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeRequestEmailCodeRequest returns a decoder for requests sent to the
// auth RequestEmailCode endpoint.
func DecodeRequestEmailCodeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body RequestEmailCodeRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateRequestEmailCodeRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewRequestEmailCodePayload(&body)

		return payload, nil
	}
}

// EncodeRequestSmSCodeResponse returns an encoder for responses returned by
// the auth RequestSmSCode endpoint.
func EncodeRequestSmSCodeResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*authviews.Success)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
		enc := encoder(ctx, w)
		body := NewRequestSmSCodeResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeRequestSmSCodeRequest returns a decoder for requests sent to the auth
// RequestSmSCode endpoint.
func DecodeRequestSmSCodeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body RequestSmSCodeRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateRequestSmSCodeRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewRequestSmSCodePayload(&body)

		return payload, nil
	}
}

// marshalAuthSessionToSessionResponseBody builds a value of type
// *SessionResponseBody from a value of type *auth.Session.
func marshalAuthSessionToSessionResponseBody(v *auth.Session) *SessionResponseBody {
	if v == nil {
		return nil
	}
	res := &SessionResponseBody{}
	if v.User != nil {
		res.User = marshalAuthUserToUserResponseBody(v.User)
	}
	if v.Credentials != nil {
		res.Credentials = marshalAuthCredentialsToCredentialsResponseBody(v.Credentials)
	}

	return res
}

// marshalAuthUserToUserResponseBody builds a value of type *UserResponseBody
// from a value of type *auth.User.
func marshalAuthUserToUserResponseBody(v *auth.User) *UserResponseBody {
	res := &UserResponseBody{
		ID:        v.ID,
		Username:  v.Username,
		Nickname:  v.Nickname,
		Mobile:    v.Mobile,
		Email:     v.Email,
		Type:      v.Type,
		IsActive:  v.IsActive,
		LoginTime: v.LoginTime,
	}

	return res
}

// marshalAuthCredentialsToCredentialsResponseBody builds a value of type
// *CredentialsResponseBody from a value of type *auth.Credentials.
func marshalAuthCredentialsToCredentialsResponseBody(v *auth.Credentials) *CredentialsResponseBody {
	res := &CredentialsResponseBody{
		Token:     v.Token,
		ExpiresIn: v.ExpiresIn,
	}

	return res
}

// marshalAuthviewsUserViewToUserResponseBodySimple builds a value of type
// *UserResponseBodySimple from a value of type *authviews.UserView.
func marshalAuthviewsUserViewToUserResponseBodySimple(v *authviews.UserView) *UserResponseBodySimple {
	if v == nil {
		return nil
	}
	res := &UserResponseBodySimple{
		ID:       *v.ID,
		Nickname: *v.Nickname,
		Mobile:   *v.Mobile,
		Email:    *v.Email,
		Type:     *v.Type,
		IsActive: *v.IsActive,
	}

	return res
}

// marshalAuthviewsUserViewToUserResponseBody builds a value of type
// *UserResponseBody from a value of type *authviews.UserView.
func marshalAuthviewsUserViewToUserResponseBody(v *authviews.UserView) *UserResponseBody {
	res := &UserResponseBody{
		ID:       *v.ID,
		Username: *v.Username,
		Nickname: *v.Nickname,
		Mobile:   *v.Mobile,
		Email:    *v.Email,
		Type:     *v.Type,
		IsActive: *v.IsActive,
	}

	return res
}

// marshalAuthviewsCredentialsViewToCredentialsResponseBody builds a value of
// type *CredentialsResponseBody from a value of type
// *authviews.CredentialsView.
func marshalAuthviewsCredentialsViewToCredentialsResponseBody(v *authviews.CredentialsView) *CredentialsResponseBody {
	res := &CredentialsResponseBody{
		Token:     *v.Token,
		ExpiresIn: *v.ExpiresIn,
	}

	return res
}
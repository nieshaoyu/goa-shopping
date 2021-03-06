// Code generated by goa v3.1.1, DO NOT EDIT.
//
// auth HTTP client CLI support package
//
// Command:
// $ goa gen goa-shopping/design

package client

import (
	"encoding/json"
	"fmt"
	auth "goa-shopping/gen/auth"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// BuildLoginPayload builds the payload for the auth Login endpoint from CLI
// flags.
func BuildLoginPayload(authLoginBody string) (*auth.LoginPayload, error) {
	var err error
	var body LoginRequestBody
	{
		err = json.Unmarshal([]byte(authLoginBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"password\": \"password\",\n      \"phone\": \"18396666666\"\n   }'")
		}
		if utf8.RuneCountInString(body.Phone) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.phone", body.Phone, utf8.RuneCountInString(body.Phone), 1, true))
		}
		if utf8.RuneCountInString(body.Phone) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.phone", body.Phone, utf8.RuneCountInString(body.Phone), 20, false))
		}
		if utf8.RuneCountInString(body.Password) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", body.Password, utf8.RuneCountInString(body.Password), 1, true))
		}
		if utf8.RuneCountInString(body.Password) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", body.Password, utf8.RuneCountInString(body.Password), 128, false))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &auth.LoginPayload{
		Phone:    body.Phone,
		Password: body.Password,
	}

	return v, nil
}

// BuildGetUserInfoPayload builds the payload for the auth GetUserInfo endpoint
// from CLI flags.
func BuildGetUserInfoPayload(authGetUserInfoJWTToken string) (*auth.GetUserInfoPayload, error) {
	var jwtToken *string
	{
		if authGetUserInfoJWTToken != "" {
			jwtToken = &authGetUserInfoJWTToken
		}
	}
	v := &auth.GetUserInfoPayload{}
	v.JWTToken = jwtToken

	return v, nil
}

// BuildSignUpPayload builds the payload for the auth SignUp endpoint from CLI
// flags.
func BuildSignUpPayload(authSignUpBody string) (*auth.SignUpPayload, error) {
	var err error
	var body SignUpRequestBody
	{
		err = json.Unmarshal([]byte(authSignUpBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"email\": \"ddu\",\n      \"nickname\": \"2zc\",\n      \"password\": \"nac\",\n      \"phone\": \"18399999999\",\n      \"verify_code\": \"hrx\"\n   }'")
		}
		if body.Phone != nil {
			if utf8.RuneCountInString(*body.Phone) > 11 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("body.phone", *body.Phone, utf8.RuneCountInString(*body.Phone), 11, false))
			}
		}
		if body.Email != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
		}
		if body.Email != nil {
			if utf8.RuneCountInString(*body.Email) > 120 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("body.email", *body.Email, utf8.RuneCountInString(*body.Email), 120, false))
			}
		}
		if utf8.RuneCountInString(body.Nickname) > 120 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.nickname", body.Nickname, utf8.RuneCountInString(body.Nickname), 120, false))
		}
		if utf8.RuneCountInString(body.Password) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", body.Password, utf8.RuneCountInString(body.Password), 6, true))
		}
		if utf8.RuneCountInString(body.VerifyCode) < 4 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.verify_code", body.VerifyCode, utf8.RuneCountInString(body.VerifyCode), 4, true))
		}
		if utf8.RuneCountInString(body.VerifyCode) > 12 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.verify_code", body.VerifyCode, utf8.RuneCountInString(body.VerifyCode), 12, false))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &auth.SignUpPayload{
		Phone:      body.Phone,
		Email:      body.Email,
		Nickname:   body.Nickname,
		Password:   body.Password,
		VerifyCode: body.VerifyCode,
	}

	return v, nil
}

// BuildSetPasswordPayload builds the payload for the auth SetPassword endpoint
// from CLI flags.
func BuildSetPasswordPayload(authSetPasswordBody string) (*auth.SetPasswordPayload, error) {
	var err error
	var body SetPasswordRequestBody
	{
		err = json.Unmarshal([]byte(authSetPasswordBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"new_password\": \"abc123455\",\n      \"username\": \"13839003900\",\n      \"verify_code\": \"888888\"\n   }'")
		}
		if utf8.RuneCountInString(body.Username) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.username", body.Username, utf8.RuneCountInString(body.Username), 128, false))
		}
		if utf8.RuneCountInString(body.NewPassword) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.new_password", body.NewPassword, utf8.RuneCountInString(body.NewPassword), 6, true))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &auth.SetPasswordPayload{
		Username:    body.Username,
		NewPassword: body.NewPassword,
		VerifyCode:  body.VerifyCode,
	}

	return v, nil
}

// BuildRequestEmailCodePayload builds the payload for the auth
// RequestEmailCode endpoint from CLI flags.
func BuildRequestEmailCodePayload(authRequestEmailCodeBody string) (*auth.RequestEmailCodePayload, error) {
	var err error
	var body RequestEmailCodeRequestBody
	{
		err = json.Unmarshal([]byte(authRequestEmailCodeBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"action\": \"register\",\n      \"email\": \"user@example.com\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))

		if !(body.Action == "register" || body.Action == "reset_password" || body.Action == "update_email") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.action", body.Action, []interface{}{"register", "reset_password", "update_email"}))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &auth.RequestEmailCodePayload{
		Email:  body.Email,
		Action: body.Action,
	}

	return v, nil
}

// BuildRequestSmSCodePayload builds the payload for the auth RequestSmSCode
// endpoint from CLI flags.
func BuildRequestSmSCodePayload(authRequestSmSCodeBody string) (*auth.RequestSmSCodePayload, error) {
	var err error
	var body RequestSmSCodeRequestBody
	{
		err = json.Unmarshal([]byte(authRequestSmSCodeBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"action\": \"register\",\n      \"mobile\": \"13838003800\"\n   }'")
		}
		if utf8.RuneCountInString(body.Mobile) < 11 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.mobile", body.Mobile, utf8.RuneCountInString(body.Mobile), 11, true))
		}
		if utf8.RuneCountInString(body.Mobile) > 11 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.mobile", body.Mobile, utf8.RuneCountInString(body.Mobile), 11, false))
		}
		if !(body.Action == "register" || body.Action == "reset_password" || body.Action == "update_mobile") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.action", body.Action, []interface{}{"register", "reset_password", "update_mobile"}))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &auth.RequestSmSCodePayload{
		Mobile: body.Mobile,
		Action: body.Action,
	}

	return v, nil
}

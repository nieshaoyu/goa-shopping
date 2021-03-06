// Code generated by goa v3.1.1, DO NOT EDIT.
//
// auth gRPC client types
//
// Command:
// $ goa gen goa-shopping/design

package client

import (
	auth "goa-shopping/gen/auth"
	authviews "goa-shopping/gen/auth/views"
	authpb "goa-shopping/gen/grpc/auth/pb"
)

// NewRequestEmailCodeRequest builds the gRPC request type from the payload of
// the "RequestEmailCode" endpoint of the "auth" service.
func NewRequestEmailCodeRequest(payload *auth.RequestEmailCodePayload) *authpb.RequestEmailCodeRequest {
	message := &authpb.RequestEmailCodeRequest{
		Email:  payload.Email,
		Action: payload.Action,
	}
	return message
}

// NewRequestEmailCodeResult builds the result type of the "RequestEmailCode"
// endpoint of the "auth" service from the gRPC response type.
func NewRequestEmailCodeResult(message *authpb.RequestEmailCodeResponse) *authviews.SuccessView {
	result := &authviews.SuccessView{
		OK: &message.Ok,
	}
	return result
}

// NewRequestSmSCodeRequest builds the gRPC request type from the payload of
// the "RequestSmSCode" endpoint of the "auth" service.
func NewRequestSmSCodeRequest(payload *auth.RequestSmSCodePayload) *authpb.RequestSmSCodeRequest {
	message := &authpb.RequestSmSCodeRequest{
		Mobile: payload.Mobile,
		Action: payload.Action,
	}
	return message
}

// NewRequestSmSCodeResult builds the result type of the "RequestSmSCode"
// endpoint of the "auth" service from the gRPC response type.
func NewRequestSmSCodeResult(message *authpb.RequestSmSCodeResponse) *authviews.SuccessView {
	result := &authviews.SuccessView{
		OK: &message.Ok,
	}
	return result
}

package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/zaplogger"
)

var _ = API("shopping", func() {
	Title("Shopping")
	Description("在线商城服务")
	Version("1.0")

	HTTP(func() {
		Path("/")
		Consumes("application/json")
		Produces("application/json")
	})

	Server("shopping", func() {
		Host("development", func() {
			URI("http://localhost:8080")
			URI("grpc://localhost:8082")
		})
	})
})

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`JWT 认证`)
	Scope("role:user", "普通用户")
	Scope("role:admin", "管理员")
})

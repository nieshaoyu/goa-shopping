package design

import (
	. "goa.design/goa/v3/dsl"
)

var (
	ExampleJwt  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
	ExampleUUID = "6ff74de6-ece0-4fbd-9a4c-b9e95dc7a374"
)

var User = ResultType("User", func() {
	Description("用户")
	ContentType("application/json")

	Attributes(func() {
		Attribute("id", Int, "ID")
		Attribute("username", String, "用户名")
		Attribute("nickname", String, "昵称")
		Attribute("mobile", String, "手机号")
		Attribute("email", String, "邮箱")
		Attribute("type", Int, "类型：1 用户， 2 管理员")
		Attribute("isActive", Boolean, "是否可用")
		Attribute("loginTime", String, "登录时间")

		Required("id", "username", "nickname", "mobile", "email", "type", "isActive", "loginTime")
	})

	View("default", func() {
		Attribute("id")
		Attribute("username")
		Attribute("nickname")
		Attribute("mobile")
		Attribute("email")
		Attribute("type")
		Attribute("isActive")
	})

	View("simple", func() {
		Attribute("id")
		Attribute("nickname")
		Attribute("mobile")
		Attribute("email")
		Attribute("type")
		Attribute("isActive")
	})
})

var Credentials = Type("Credentials", func() {
	Field(1, "token", String, "JWT token", func() {
		Example(ExampleJwt)
	})
	Field(7, "expires_in", Int, "有效时长（秒）：生成之后x秒内有效", func() {
		Example(25200)
	})
	Required("token", "expires_in")
})

var Session = ResultType("Session", func() {
	Description("会话")
	ContentType("application/json")
	Attributes(func() {
		Field(1, "user", User)
		Field(2, "credentials", Credentials)
		Required("user", "credentials")
	})

	View("default", func() {
		Attribute("user")
		Attribute("credentials")
	})
})

var SuccessResult = ResultType("application/vnd.success_result", func() {
	Description("成功")
	TypeName("success")
	ContentType("application/json")

	Attributes(func() {
		Field(1, "ok", Boolean, "success", func() {
			Example(true)
		})
		Required("ok")
	})

	View("default", func() {
		Attribute("ok")
	})
})

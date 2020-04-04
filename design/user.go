package design

import (
	. "goa.design/goa/v3/dsl"
)

var UserInfo = ResultType("UserInfo", func() {
	TypeName("UserInfo")
	ContentType("application/json")

	Attributes(func() {
		Attribute("errcode", Int, "错误码(0: 成功; 大于0: 失败)", func() {
			Minimum(0)
			Maximum(999999)
			Example(0)
		})
		Attribute("errmsg", String, "错误消息", func() {
			Example("")
		})
		Attribute("data", User, "用户信息", func() {
			View("simple")
		})
		Required("errcode", "errmsg")
	})

	View("default", func() {
		Attribute("errcode")
		Attribute("errmsg")
		Attribute("data")
	})
})

var _ = Service("auth", func() {
	Description("用户服务")

	Error("internal_server_error", ErrorResult, func() {
		Fault()
	})
	Error("bad_request", ErrorResult)

	Method("Login", func() {
		Description("使用手机号或者邮箱 + 密码登录")
		Meta("swagger:summary", "使用手机号或者邮箱 + 密码登录")
		Payload(func() {
			Attribute("phone", String, "账号(手机号或者邮箱)", func() {
				Example("18396666666")
				MinLength(1)
				MaxLength(20)
			})
			Attribute("password", String, "密码", func() {
				Example("password")
				MinLength(1)
				MaxLength(128)
			})
			Required("phone", "password")
		})

		Result(func() {
			Attribute("errcode", Int, "错误码", func() {
				Minimum(0)
				Maximum(999999)
				Example(0)
			})
			Attribute("errmsg", String, "错误消息", func() {
				Example("")
			})
			Attribute("data", Session)
			Required("errcode", "errmsg")
		})

		HTTP(func() {
			POST("/login")
			Response(StatusOK)
		})
	})

	Method("GetUserInfo", func() {
		Description("获取用户信息")
		Meta("swagger:summary", "获取用户信息")

		Security(JWTAuth, func() {
			Scope("role:user")
		})

		Payload(func() {
			Token("jwtToken", String, "jwt token")
		})

		Result(UserInfo)
		HTTP(func() {
			POST("/get_user_info")
			Response(StatusOK)
		})
	})

	Method("SignUp", func() {
		Meta("swagger:summary", "用户注册")
		Description(`用户注册使用手机/邮箱+验证码注册`)
		Payload(func() {
			Field(1, "phone", String, "Phone", func() {
				Example("18399999999")
				MaxLength(11)
			})
			Field(2, "email", String, "Email", func() {
				Format(FormatEmail)
				MaxLength(120)
			})
			Field(3, "nickname", String, "昵称", func() {
				MaxLength(120)
			})
			Field(4, "password", String, "密码", func() {
				MinLength(6)
			})
			Field(5, "verify_code", String, "验证码", func() {
				MaxLength(12)
				MinLength(4)
			})
			Required("nickname", "password", "verify_code")
		})

		Result(Session)

		HTTP(func() {
			POST("/sign_up")
			Response(StatusOK)
			Response("bad_request", StatusBadRequest)
		})
	})

	Method("SetPassword", func() {
		Meta("swagger:summary", "重置密码")
		Description("重置密码")

		Payload(func() {
			Field(1, "username", String, "用户名（邮箱/手机号）", func() {
				MaxLength(128)
				Example("13839003900")
			})
			Field(2, "new_password", String, "新密码", func() {
				MinLength(6)
				Example("abc123455")
			})
			Field(3, "verify_code", String, "验证码", func() {
				Example("888888")
			})
			Required("username", "new_password", "verify_code")
		})

		Result(SuccessResult)

		HTTP(func() {
			POST("/reset_password")
			Response(StatusOK)
		})
	})

	Method("RequestEmailCode", func() {
		Meta("swagger:summary", "请求电子邮箱验证码")
		Description("请求电子邮箱验证码")
		Payload(func() {
			Field(1, "email", String, "电子邮箱", func() {
				Format(FormatEmail)
				Example("user@example.com")
			})
			Field(2, "action", String, "操作", func() {
				Enum("register", "reset_password", "update_email")
				Example("register")
			})
			Required("email", "action")
		})

		Result(SuccessResult)

		HTTP(func() {
			POST("/request_email_code")
			Response(StatusOK)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("RequestSmSCode", func() {
		Meta("swagger:summary", "请求短信验证码")
		Description("请求短信验证码")
		Payload(func() {
			Field(1, "mobile", String, "手机号码", func() {
				MaxLength(11)
				MinLength(11)
				Example("13838003800")
			})
			Field(2, "action", String, "操作", func() {
				Enum("register", "reset_password", "update_mobile")
				Example("register")
			})
			Required("mobile", "action")
		})

		Result(SuccessResult)

		HTTP(func() {
			POST("/request_sms_code")
			Response(StatusOK)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})

})

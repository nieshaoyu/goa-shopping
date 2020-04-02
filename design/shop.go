package design

import . "goa.design/goa/v3/dsl"

var _ = Service("shop", func() {
	Description("在线商城")

	Error("internal_server_error", ErrorResult, func() {
		Fault()
	})
	Error("bad_request", ErrorResult)

	HTTP(func() {
		Path("/shop")
		Response("internal_server_error", StatusInternalServerError)
		Response("bad_request", StatusBadRequest)
	})

	Method("GoodsList", func() {
		Description("商品列表")
		Meta("swagger:summary", "分组列表")

		Security(JWTAuth, func() {
			Scope("role:user")
			Scope("role:admin")
		})

		Payload(func() {
			Token("jwtToken", String, func() {
				Description("JWT used for authentication")
				Example(ExampleJwt)
			})
			Field(1, "cursor", Int, "cursor of page", func() {
				Example(0)
			})
			Field(2, "limit", Int, "limit of items", func() {
				Example(20)
			})
			Required("cursor", "limit")
		})

		Result(func() {
			Field(1, "items", ArrayOf(Goods))
			Field(2, "nextCursor", Int, "下一页游标", func() {
				Example(100)
			})
			Field(3, "total", Int, "总记录数")
			Required("items", "nextCursor", "total")
		})

		HTTP(func() {
			GET("/")
			Params(func() {
				Param("cursor")
				Param("limit")
			})
			Response(StatusOK)
		})

	})

	Method("GetGoods", func() {
		Description("商品详情")
		Meta("swagger:summary", "商品详情")

		Security(JWTAuth, func() {
			Scope("role:user")
			Scope("role:admin")
		})

		Payload(func() {
			Token("jwtToken", String, func() {
				Description("JWT used for authentication")
				Example(ExampleJwt)
			})
			Field(1, "goodsId", String, func() {
				MinLength(1)
			})
			Required("jwtToken", "goodsId")
		})

		Result(GoodsSpu)

		HTTP(func() {
			GET("/{goodsId}")
			Response(StatusOK)
		})
	})

	Method("UpdateGoods", func() {
		Description("更新商品")
		Meta("swagger:summary", "更新商品")

		Security(JWTAuth, func() {
			Scope("role:admin")
		})

		Payload(func() {
			Token("jwtToken", String, func() {
				Description("JWT used for authentication")
				Example(ExampleJwt)
			})
			Field(1, "goodsId", String, func() {
				MinLength(1)
			})
			Field(2, "name", String, "商品名称", func() {
				Example("水果x")
				MaxLength(20)
			})
			Field(3, "description", String, "商品描述", func() {
				Example("水果x刘海屏就很丑")
				MaxLength(128)
			})
			Required("jwtToken", "goodsId", "name", "description")
		})

		Result(GoodsSpu)

		HTTP(func() {
			PUT("/{goodsId}")
			Response(StatusOK)
		})

	})

	Files("/apidoc.html", "gen/apidoc.html")
})

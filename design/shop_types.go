package design

import . "goa.design/goa/v3/dsl"

var GoodsType = ResultType("goodsType", func() {
	Description("商品分类")
	TypeName("goodsType")
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", Int, "ID")
		Attribute("name", String, "类型名称", func() {
			Example("男装")
		})
		Attribute("description", String, "类型描述", func() {
			Example("男人的衣柜")
		})
		Required("id", "name", "description")
	})
})

var GoodsBrand = ResultType("goodsBrand", func() {
	Description("商品品牌")
	TypeName("goodsBrand")
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", Int, "ID")
		Attribute("name", String, "品牌名称", func() {
			Example("耐克")
		})
		Attribute("description", String, "品牌描述", func() {
			Example("耐克描述")
		})
		Required("id", "name", "description")
	})
})

var GoodsSku = ResultType("goodsSku", func() {
	Description("商品SKU")
	TypeName("goodsSku")
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", Int, "ID")
		Attribute("name", String, "颜色", func() {
			Example("颜色")
		})
		Attribute("goods_sku_value", ArrayOf(GoodsSkuValue), "商品SKU规格值")
		Required("id", "name")
	})
})

var GoodsSkuValue = ResultType("goodsSkuValue", func() {
	Description("商品SKU规格值")
	TypeName("goodsSkuValue")
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", Int, "ID")
		Attribute("name", String, "红色", func() {
			Example("红色")
		})
		Required("id", "name")
	})
})

var GoodsSpu = ResultType("goodsSpu", func() {
	Description("商品SPU")
	TypeName("goodsSpu")
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", Int, "ID")
		Attribute("name", String, "商品名称", func() {
			Example("水果机")
		})
		Attribute("description", String, "商品描述", func() {
			Example("水果机是世界上最好的手机")
		})
		Attribute("goods_type", GoodsType, "商品类型") // fixme 理论上一个商品可以属于很多类型, 但是我这里没有理论
		Attribute("goods_brand", GoodsBrand, "商品品牌")
		Attribute("goods_sku", ArrayOf(GoodsSku), "商品SKU")

		Required("id", "name", "description", "goods_type", "goods_brand")
	})
})

var Goods = ResultType("goods", func() {
	Description("商品Spu")
	TypeName("goods")
	ContentType("application/json")
	Attributes(func() {
		Attribute("errcode", Int, "错误码", func() {
			Minimum(0)
			Maximum(999999)
			Example(0)
		})
		Attribute("errmsg", String, "错误消息", func() {
			Example("")
		})
		Attribute("data", ArrayOf(GoodsSpu, func() {
			View("default")
		}))
		Required("errcode", "errmsg", "data")
	})
})

// Code generated by goa v3.1.1, DO NOT EDIT.
//
// shop HTTP server types
//
// Command:
// $ goa gen goa-shopping/design

package server

import (
	shop "goa-shopping/gen/shop"
	shopviews "goa-shopping/gen/shop/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// UpdateGoodsRequestBody is the type of the "shop" service "UpdateGoods"
// endpoint HTTP request body.
type UpdateGoodsRequestBody struct {
	// 商品名称
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 商品描述
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// GoodsListResponseBody is the type of the "shop" service "GoodsList" endpoint
// HTTP response body.
type GoodsListResponseBody struct {
	Items []*GoodsResponseBody `form:"items" json:"items" xml:"items"`
	// 下一页游标
	NextCursor int `form:"nextCursor" json:"nextCursor" xml:"nextCursor"`
	// 总记录数
	Total int `form:"total" json:"total" xml:"total"`
}

// GetGoodsResponseBody is the type of the "shop" service "GetGoods" endpoint
// HTTP response body.
type GetGoodsResponseBody struct {
	// ID
	ID int `form:"id" json:"id" xml:"id"`
	// 商品名称
	Name string `form:"name" json:"name" xml:"name"`
	// 商品描述
	Description string `form:"description" json:"description" xml:"description"`
	// 商品类型
	GoodsType *GoodsTypeResponseBody `form:"goods_type" json:"goods_type" xml:"goods_type"`
	// 商品品牌
	GoodsBrand *GoodsBrandResponseBody `form:"goods_brand" json:"goods_brand" xml:"goods_brand"`
	// 商品SKU
	GoodsSku []*GoodsSkuResponseBody `form:"goods_sku,omitempty" json:"goods_sku,omitempty" xml:"goods_sku,omitempty"`
}

// UpdateGoodsResponseBody is the type of the "shop" service "UpdateGoods"
// endpoint HTTP response body.
type UpdateGoodsResponseBody struct {
	// ID
	ID int `form:"id" json:"id" xml:"id"`
	// 商品名称
	Name string `form:"name" json:"name" xml:"name"`
	// 商品描述
	Description string `form:"description" json:"description" xml:"description"`
	// 商品类型
	GoodsType *GoodsTypeResponseBody `form:"goods_type" json:"goods_type" xml:"goods_type"`
	// 商品品牌
	GoodsBrand *GoodsBrandResponseBody `form:"goods_brand" json:"goods_brand" xml:"goods_brand"`
	// 商品SKU
	GoodsSku []*GoodsSkuResponseBody `form:"goods_sku,omitempty" json:"goods_sku,omitempty" xml:"goods_sku,omitempty"`
}

// GoodsListInternalServerErrorResponseBody is the type of the "shop" service
// "GoodsList" endpoint HTTP response body for the "internal_server_error"
// error.
type GoodsListInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GoodsListBadRequestResponseBody is the type of the "shop" service
// "GoodsList" endpoint HTTP response body for the "bad_request" error.
type GoodsListBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetGoodsInternalServerErrorResponseBody is the type of the "shop" service
// "GetGoods" endpoint HTTP response body for the "internal_server_error" error.
type GetGoodsInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetGoodsBadRequestResponseBody is the type of the "shop" service "GetGoods"
// endpoint HTTP response body for the "bad_request" error.
type GetGoodsBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateGoodsInternalServerErrorResponseBody is the type of the "shop" service
// "UpdateGoods" endpoint HTTP response body for the "internal_server_error"
// error.
type UpdateGoodsInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateGoodsBadRequestResponseBody is the type of the "shop" service
// "UpdateGoods" endpoint HTTP response body for the "bad_request" error.
type UpdateGoodsBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GoodsResponseBody is used to define fields on response body types.
type GoodsResponseBody struct {
	// 错误码
	Errcode int `form:"errcode" json:"errcode" xml:"errcode"`
	// 错误消息
	Errmsg string                  `form:"errmsg" json:"errmsg" xml:"errmsg"`
	Data   []*GoodsSpuResponseBody `form:"data" json:"data" xml:"data"`
}

// GoodsSpuResponseBody is used to define fields on response body types.
type GoodsSpuResponseBody struct {
	// ID
	ID int `form:"id" json:"id" xml:"id"`
	// 商品名称
	Name string `form:"name" json:"name" xml:"name"`
	// 商品描述
	Description string `form:"description" json:"description" xml:"description"`
	// 商品类型
	GoodsType *GoodsTypeResponseBody `form:"goods_type" json:"goods_type" xml:"goods_type"`
	// 商品品牌
	GoodsBrand *GoodsBrandResponseBody `form:"goods_brand" json:"goods_brand" xml:"goods_brand"`
	// 商品SKU
	GoodsSku []*GoodsSkuResponseBody `form:"goods_sku,omitempty" json:"goods_sku,omitempty" xml:"goods_sku,omitempty"`
}

// GoodsTypeResponseBody is used to define fields on response body types.
type GoodsTypeResponseBody struct {
	// ID
	ID int `form:"id" json:"id" xml:"id"`
	// 类型名称
	Name string `form:"name" json:"name" xml:"name"`
	// 类型描述
	Description string `form:"description" json:"description" xml:"description"`
}

// GoodsBrandResponseBody is used to define fields on response body types.
type GoodsBrandResponseBody struct {
	// ID
	ID int `form:"id" json:"id" xml:"id"`
	// 品牌名称
	Name string `form:"name" json:"name" xml:"name"`
	// 品牌描述
	Description string `form:"description" json:"description" xml:"description"`
}

// GoodsSkuResponseBody is used to define fields on response body types.
type GoodsSkuResponseBody struct {
	// ID
	ID int `form:"id" json:"id" xml:"id"`
	// 颜色
	Name string `form:"name" json:"name" xml:"name"`
	// 商品SKU规格值
	GoodsSkuValue []*GoodsSkuValueResponseBody `form:"goods_sku_value,omitempty" json:"goods_sku_value,omitempty" xml:"goods_sku_value,omitempty"`
}

// GoodsSkuValueResponseBody is used to define fields on response body types.
type GoodsSkuValueResponseBody struct {
	// ID
	ID int `form:"id" json:"id" xml:"id"`
	// 红色
	Name string `form:"name" json:"name" xml:"name"`
}

// NewGoodsListResponseBody builds the HTTP response body from the result of
// the "GoodsList" endpoint of the "shop" service.
func NewGoodsListResponseBody(res *shop.GoodsListResult) *GoodsListResponseBody {
	body := &GoodsListResponseBody{
		NextCursor: res.NextCursor,
		Total:      res.Total,
	}
	if res.Items != nil {
		body.Items = make([]*GoodsResponseBody, len(res.Items))
		for i, val := range res.Items {
			body.Items[i] = marshalShopGoodsToGoodsResponseBody(val)
		}
	}
	return body
}

// NewGetGoodsResponseBody builds the HTTP response body from the result of the
// "GetGoods" endpoint of the "shop" service.
func NewGetGoodsResponseBody(res *shopviews.GoodsSpuView) *GetGoodsResponseBody {
	body := &GetGoodsResponseBody{
		ID:          *res.ID,
		Name:        *res.Name,
		Description: *res.Description,
	}
	if res.GoodsType != nil {
		body.GoodsType = marshalShopviewsGoodsTypeViewToGoodsTypeResponseBody(res.GoodsType)
	}
	if res.GoodsBrand != nil {
		body.GoodsBrand = marshalShopviewsGoodsBrandViewToGoodsBrandResponseBody(res.GoodsBrand)
	}
	if res.GoodsSku != nil {
		body.GoodsSku = make([]*GoodsSkuResponseBody, len(res.GoodsSku))
		for i, val := range res.GoodsSku {
			body.GoodsSku[i] = marshalShopviewsGoodsSkuViewToGoodsSkuResponseBody(val)
		}
	}
	return body
}

// NewUpdateGoodsResponseBody builds the HTTP response body from the result of
// the "UpdateGoods" endpoint of the "shop" service.
func NewUpdateGoodsResponseBody(res *shopviews.GoodsSpuView) *UpdateGoodsResponseBody {
	body := &UpdateGoodsResponseBody{
		ID:          *res.ID,
		Name:        *res.Name,
		Description: *res.Description,
	}
	if res.GoodsType != nil {
		body.GoodsType = marshalShopviewsGoodsTypeViewToGoodsTypeResponseBody(res.GoodsType)
	}
	if res.GoodsBrand != nil {
		body.GoodsBrand = marshalShopviewsGoodsBrandViewToGoodsBrandResponseBody(res.GoodsBrand)
	}
	if res.GoodsSku != nil {
		body.GoodsSku = make([]*GoodsSkuResponseBody, len(res.GoodsSku))
		for i, val := range res.GoodsSku {
			body.GoodsSku[i] = marshalShopviewsGoodsSkuViewToGoodsSkuResponseBody(val)
		}
	}
	return body
}

// NewGoodsListInternalServerErrorResponseBody builds the HTTP response body
// from the result of the "GoodsList" endpoint of the "shop" service.
func NewGoodsListInternalServerErrorResponseBody(res *goa.ServiceError) *GoodsListInternalServerErrorResponseBody {
	body := &GoodsListInternalServerErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGoodsListBadRequestResponseBody builds the HTTP response body from the
// result of the "GoodsList" endpoint of the "shop" service.
func NewGoodsListBadRequestResponseBody(res *goa.ServiceError) *GoodsListBadRequestResponseBody {
	body := &GoodsListBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetGoodsInternalServerErrorResponseBody builds the HTTP response body
// from the result of the "GetGoods" endpoint of the "shop" service.
func NewGetGoodsInternalServerErrorResponseBody(res *goa.ServiceError) *GetGoodsInternalServerErrorResponseBody {
	body := &GetGoodsInternalServerErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetGoodsBadRequestResponseBody builds the HTTP response body from the
// result of the "GetGoods" endpoint of the "shop" service.
func NewGetGoodsBadRequestResponseBody(res *goa.ServiceError) *GetGoodsBadRequestResponseBody {
	body := &GetGoodsBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateGoodsInternalServerErrorResponseBody builds the HTTP response body
// from the result of the "UpdateGoods" endpoint of the "shop" service.
func NewUpdateGoodsInternalServerErrorResponseBody(res *goa.ServiceError) *UpdateGoodsInternalServerErrorResponseBody {
	body := &UpdateGoodsInternalServerErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateGoodsBadRequestResponseBody builds the HTTP response body from the
// result of the "UpdateGoods" endpoint of the "shop" service.
func NewUpdateGoodsBadRequestResponseBody(res *goa.ServiceError) *UpdateGoodsBadRequestResponseBody {
	body := &UpdateGoodsBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGoodsListPayload builds a shop service GoodsList endpoint payload.
func NewGoodsListPayload(cursor int, limit int, jwtToken *string) *shop.GoodsListPayload {
	v := &shop.GoodsListPayload{}
	v.Cursor = cursor
	v.Limit = limit
	v.JWTToken = jwtToken

	return v
}

// NewGetGoodsPayload builds a shop service GetGoods endpoint payload.
func NewGetGoodsPayload(goodsID string, jwtToken string) *shop.GetGoodsPayload {
	v := &shop.GetGoodsPayload{}
	v.GoodsID = goodsID
	v.JWTToken = jwtToken

	return v
}

// NewUpdateGoodsPayload builds a shop service UpdateGoods endpoint payload.
func NewUpdateGoodsPayload(body *UpdateGoodsRequestBody, goodsID string, jwtToken string) *shop.UpdateGoodsPayload {
	v := &shop.UpdateGoodsPayload{
		Name:        *body.Name,
		Description: *body.Description,
	}
	v.GoodsID = goodsID
	v.JWTToken = jwtToken

	return v
}

// ValidateUpdateGoodsRequestBody runs the validations defined on
// UpdateGoodsRequestBody
func ValidateUpdateGoodsRequestBody(body *UpdateGoodsRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 20, false))
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 128, false))
		}
	}
	return
}
// Code generated by goa v3.1.1, DO NOT EDIT.
//
// shop service
//
// Command:
// $ goa gen goa-shopping/design

package shop

import (
	"context"
	shopviews "goa-shopping/gen/shop/views"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// 在线商城
type Service interface {
	// 商品列表
	GoodsList(context.Context, *GoodsListPayload) (res *GoodsListResult, err error)
	// 商品详情
	GetGoods(context.Context, *GetGoodsPayload) (res *GoodsSpu, err error)
	// 更新商品
	UpdateGoods(context.Context, *UpdateGoodsPayload) (res *GoodsSpu, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "shop"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"GoodsList", "GetGoods", "UpdateGoods"}

// GoodsListPayload is the payload type of the shop service GoodsList method.
type GoodsListPayload struct {
	// JWT used for authentication
	JWTToken *string `json:"jwt_token"`
	// cursor of page
	Cursor int `json:"cursor"`
	// limit of items
	Limit int `json:"limit"`
}

// GoodsListResult is the result type of the shop service GoodsList method.
type GoodsListResult struct {
	Items []*Goods `json:"items"`
	// 下一页游标
	NextCursor int `json:"next_cursor"`
	// 总记录数
	Total int `json:"total"`
}

// GetGoodsPayload is the payload type of the shop service GetGoods method.
type GetGoodsPayload struct {
	// JWT used for authentication
	JWTToken string `json:"jwt_token"`
	GoodsID  string `json:"goods_id"`
}

// GoodsSpu is the result type of the shop service GetGoods method.
type GoodsSpu struct {
	// ID
	ID int `json:"id"`
	// 商品名称
	Name string `json:"name"`
	// 商品描述
	Description string `json:"description"`
	// 商品类型
	GoodsType *GoodsType `json:"goods_type"`
	// 商品品牌
	GoodsBrand *GoodsBrand `json:"goods_brand"`
	// 商品SKU
	GoodsSku []*GoodsSku `json:"goods_sku"`
}

// UpdateGoodsPayload is the payload type of the shop service UpdateGoods
// method.
type UpdateGoodsPayload struct {
	// JWT used for authentication
	JWTToken string `json:"jwt_token"`
	GoodsID  string `json:"goods_id"`
	// 商品名称
	Name string `json:"name"`
	// 商品描述
	Description string `json:"description"`
}

// 商品Spu
type Goods struct {
	// 错误码
	Errcode int `json:"errcode"`
	// 错误消息
	Errmsg string      `json:"errmsg"`
	Data   []*GoodsSpu `json:"data"`
}

// 商品分类
type GoodsType struct {
	// ID
	ID int `json:"id"`
	// 类型名称
	Name string `json:"name"`
	// 类型描述
	Description string `json:"description"`
}

// 商品品牌
type GoodsBrand struct {
	// ID
	ID int `json:"id"`
	// 品牌名称
	Name string `json:"name"`
	// 品牌描述
	Description string `json:"description"`
}

// 商品SKU
type GoodsSku struct {
	// ID
	ID int `json:"id"`
	// 颜色
	Name string `json:"name"`
	// 商品SKU规格值
	GoodsSkuValue []*GoodsSkuValue `json:"goods_sku_value"`
}

// 商品SKU规格值
type GoodsSkuValue struct {
	// ID
	ID int `json:"id"`
	// 红色
	Name string `json:"name"`
}

// MakeInternalServerError builds a goa.ServiceError from an error.
func MakeInternalServerError(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "internal_server_error",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
		Fault:   true,
	}
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "bad_request",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewGoodsSpu initializes result type GoodsSpu from viewed result type
// GoodsSpu.
func NewGoodsSpu(vres *shopviews.GoodsSpu) *GoodsSpu {
	return newGoodsSpu(vres.Projected)
}

// NewViewedGoodsSpu initializes viewed result type GoodsSpu from result type
// GoodsSpu using the given view.
func NewViewedGoodsSpu(res *GoodsSpu, view string) *shopviews.GoodsSpu {
	p := newGoodsSpuView(res)
	return &shopviews.GoodsSpu{Projected: p, View: "default"}
}

// newGoodsSpu converts projected type GoodsSpu to service type GoodsSpu.
func newGoodsSpu(vres *shopviews.GoodsSpuView) *GoodsSpu {
	res := &GoodsSpu{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Description != nil {
		res.Description = *vres.Description
	}
	if vres.GoodsBrand != nil {
		res.GoodsBrand = transformShopviewsGoodsBrandViewToGoodsBrand(vres.GoodsBrand)
	}
	if vres.GoodsSku != nil {
		res.GoodsSku = make([]*GoodsSku, len(vres.GoodsSku))
		for i, val := range vres.GoodsSku {
			res.GoodsSku[i] = transformShopviewsGoodsSkuViewToGoodsSku(val)
		}
	}
	if vres.GoodsType != nil {
		res.GoodsType = newGoodsType(vres.GoodsType)
	}
	return res
}

// newGoodsSpuView projects result type GoodsSpu to projected type GoodsSpuView
// using the "default" view.
func newGoodsSpuView(res *GoodsSpu) *shopviews.GoodsSpuView {
	vres := &shopviews.GoodsSpuView{
		ID:          &res.ID,
		Name:        &res.Name,
		Description: &res.Description,
	}
	if res.GoodsBrand != nil {
		vres.GoodsBrand = transformGoodsBrandToShopviewsGoodsBrandView(res.GoodsBrand)
	}
	if res.GoodsSku != nil {
		vres.GoodsSku = make([]*shopviews.GoodsSkuView, len(res.GoodsSku))
		for i, val := range res.GoodsSku {
			vres.GoodsSku[i] = transformGoodsSkuToShopviewsGoodsSkuView(val)
		}
	}
	if res.GoodsType != nil {
		vres.GoodsType = newGoodsTypeView(res.GoodsType)
	}
	return vres
}

// newGoodsType converts projected type GoodsType to service type GoodsType.
func newGoodsType(vres *shopviews.GoodsTypeView) *GoodsType {
	res := &GoodsType{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Description != nil {
		res.Description = *vres.Description
	}
	return res
}

// newGoodsTypeView projects result type GoodsType to projected type
// GoodsTypeView using the "default" view.
func newGoodsTypeView(res *GoodsType) *shopviews.GoodsTypeView {
	vres := &shopviews.GoodsTypeView{
		ID:          &res.ID,
		Name:        &res.Name,
		Description: &res.Description,
	}
	return vres
}

// newGoodsBrand converts projected type GoodsBrand to service type GoodsBrand.
func newGoodsBrand(vres *shopviews.GoodsBrandView) *GoodsBrand {
	res := &GoodsBrand{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Description != nil {
		res.Description = *vres.Description
	}
	return res
}

// newGoodsBrandView projects result type GoodsBrand to projected type
// GoodsBrandView using the "default" view.
func newGoodsBrandView(res *GoodsBrand) *shopviews.GoodsBrandView {
	vres := &shopviews.GoodsBrandView{
		ID:          &res.ID,
		Name:        &res.Name,
		Description: &res.Description,
	}
	return vres
}

// newGoodsSku converts projected type GoodsSku to service type GoodsSku.
func newGoodsSku(vres *shopviews.GoodsSkuView) *GoodsSku {
	res := &GoodsSku{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.GoodsSkuValue != nil {
		res.GoodsSkuValue = make([]*GoodsSkuValue, len(vres.GoodsSkuValue))
		for i, val := range vres.GoodsSkuValue {
			res.GoodsSkuValue[i] = transformShopviewsGoodsSkuValueViewToGoodsSkuValue(val)
		}
	}
	return res
}

// newGoodsSkuView projects result type GoodsSku to projected type GoodsSkuView
// using the "default" view.
func newGoodsSkuView(res *GoodsSku) *shopviews.GoodsSkuView {
	vres := &shopviews.GoodsSkuView{
		ID:   &res.ID,
		Name: &res.Name,
	}
	if res.GoodsSkuValue != nil {
		vres.GoodsSkuValue = make([]*shopviews.GoodsSkuValueView, len(res.GoodsSkuValue))
		for i, val := range res.GoodsSkuValue {
			vres.GoodsSkuValue[i] = transformGoodsSkuValueToShopviewsGoodsSkuValueView(val)
		}
	}
	return vres
}

// newGoodsSkuValue converts projected type GoodsSkuValue to service type
// GoodsSkuValue.
func newGoodsSkuValue(vres *shopviews.GoodsSkuValueView) *GoodsSkuValue {
	res := &GoodsSkuValue{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	return res
}

// newGoodsSkuValueView projects result type GoodsSkuValue to projected type
// GoodsSkuValueView using the "default" view.
func newGoodsSkuValueView(res *GoodsSkuValue) *shopviews.GoodsSkuValueView {
	vres := &shopviews.GoodsSkuValueView{
		ID:   &res.ID,
		Name: &res.Name,
	}
	return vres
}

// transformShopviewsGoodsBrandViewToGoodsBrand builds a value of type
// *GoodsBrand from a value of type *shopviews.GoodsBrandView.
func transformShopviewsGoodsBrandViewToGoodsBrand(v *shopviews.GoodsBrandView) *GoodsBrand {
	if v == nil {
		return nil
	}
	res := &GoodsBrand{
		ID:          *v.ID,
		Name:        *v.Name,
		Description: *v.Description,
	}

	return res
}

// transformShopviewsGoodsSkuViewToGoodsSku builds a value of type *GoodsSku
// from a value of type *shopviews.GoodsSkuView.
func transformShopviewsGoodsSkuViewToGoodsSku(v *shopviews.GoodsSkuView) *GoodsSku {
	if v == nil {
		return nil
	}
	res := &GoodsSku{
		ID:   *v.ID,
		Name: *v.Name,
	}
	if v.GoodsSkuValue != nil {
		res.GoodsSkuValue = make([]*GoodsSkuValue, len(v.GoodsSkuValue))
		for i, val := range v.GoodsSkuValue {
			res.GoodsSkuValue[i] = transformShopviewsGoodsSkuValueViewToGoodsSkuValue(val)
		}
	}

	return res
}

// transformShopviewsGoodsSkuValueViewToGoodsSkuValue builds a value of type
// *GoodsSkuValue from a value of type *shopviews.GoodsSkuValueView.
func transformShopviewsGoodsSkuValueViewToGoodsSkuValue(v *shopviews.GoodsSkuValueView) *GoodsSkuValue {
	if v == nil {
		return nil
	}
	res := &GoodsSkuValue{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// transformGoodsBrandToShopviewsGoodsBrandView builds a value of type
// *shopviews.GoodsBrandView from a value of type *GoodsBrand.
func transformGoodsBrandToShopviewsGoodsBrandView(v *GoodsBrand) *shopviews.GoodsBrandView {
	res := &shopviews.GoodsBrandView{
		ID:          &v.ID,
		Name:        &v.Name,
		Description: &v.Description,
	}

	return res
}

// transformGoodsSkuToShopviewsGoodsSkuView builds a value of type
// *shopviews.GoodsSkuView from a value of type *GoodsSku.
func transformGoodsSkuToShopviewsGoodsSkuView(v *GoodsSku) *shopviews.GoodsSkuView {
	if v == nil {
		return nil
	}
	res := &shopviews.GoodsSkuView{
		ID:   &v.ID,
		Name: &v.Name,
	}
	if v.GoodsSkuValue != nil {
		res.GoodsSkuValue = make([]*shopviews.GoodsSkuValueView, len(v.GoodsSkuValue))
		for i, val := range v.GoodsSkuValue {
			res.GoodsSkuValue[i] = transformGoodsSkuValueToShopviewsGoodsSkuValueView(val)
		}
	}

	return res
}

// transformGoodsSkuValueToShopviewsGoodsSkuValueView builds a value of type
// *shopviews.GoodsSkuValueView from a value of type *GoodsSkuValue.
func transformGoodsSkuValueToShopviewsGoodsSkuValueView(v *GoodsSkuValue) *shopviews.GoodsSkuValueView {
	if v == nil {
		return nil
	}
	res := &shopviews.GoodsSkuValueView{
		ID:   &v.ID,
		Name: &v.Name,
	}

	return res
}
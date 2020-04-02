// Code generated by goa v3.1.1, DO NOT EDIT.
//
// shop HTTP client encoders and decoders
//
// Command:
// $ goa gen goa-shopping/design

package client

import (
	"bytes"
	"context"
	"fmt"
	shop "goa-shopping/gen/shop"
	shopviews "goa-shopping/gen/shop/views"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildGoodsListRequest instantiates a HTTP request object with method and
// path set to call the "shop" service "GoodsList" endpoint
func (c *Client) BuildGoodsListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GoodsListShopPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("shop", "GoodsList", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGoodsListRequest returns an encoder for requests sent to the shop
// GoodsList server.
func EncodeGoodsListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*shop.GoodsListPayload)
		if !ok {
			return goahttp.ErrInvalidType("shop", "GoodsList", "*shop.GoodsListPayload", v)
		}
		if p.JWTToken != nil {
			head := *p.JWTToken
			req.Header.Set("Authorization", head)
		}
		values := req.URL.Query()
		values.Add("cursor", fmt.Sprintf("%v", p.Cursor))
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeGoodsListResponse returns a decoder for responses returned by the shop
// GoodsList endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGoodsListResponse may return the following errors:
//	- "internal_server_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeGoodsListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GoodsListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("shop", "GoodsList", err)
			}
			err = ValidateGoodsListResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("shop", "GoodsList", err)
			}
			res := NewGoodsListResultOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body GoodsListInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("shop", "GoodsList", err)
			}
			err = ValidateGoodsListInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("shop", "GoodsList", err)
			}
			return nil, NewGoodsListInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body GoodsListBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("shop", "GoodsList", err)
			}
			err = ValidateGoodsListBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("shop", "GoodsList", err)
			}
			return nil, NewGoodsListBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("shop", "GoodsList", resp.StatusCode, string(body))
		}
	}
}

// BuildGetGoodsRequest instantiates a HTTP request object with method and path
// set to call the "shop" service "GetGoods" endpoint
func (c *Client) BuildGetGoodsRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		goodsID string
	)
	{
		p, ok := v.(*shop.GetGoodsPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("shop", "GetGoods", "*shop.GetGoodsPayload", v)
		}
		goodsID = p.GoodsID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetGoodsShopPath(goodsID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("shop", "GetGoods", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetGoodsRequest returns an encoder for requests sent to the shop
// GetGoods server.
func EncodeGetGoodsRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*shop.GetGoodsPayload)
		if !ok {
			return goahttp.ErrInvalidType("shop", "GetGoods", "*shop.GetGoodsPayload", v)
		}
		{
			head := p.JWTToken
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeGetGoodsResponse returns a decoder for responses returned by the shop
// GetGoods endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetGoodsResponse may return the following errors:
//	- "internal_server_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeGetGoodsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetGoodsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("shop", "GetGoods", err)
			}
			p := NewGetGoodsGoodsSpuOK(&body)
			view := "default"
			vres := &shopviews.GoodsSpu{Projected: p, View: view}
			if err = shopviews.ValidateGoodsSpu(vres); err != nil {
				return nil, goahttp.ErrValidationError("shop", "GetGoods", err)
			}
			res := shop.NewGoodsSpu(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body GetGoodsInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("shop", "GetGoods", err)
			}
			err = ValidateGetGoodsInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("shop", "GetGoods", err)
			}
			return nil, NewGetGoodsInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body GetGoodsBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("shop", "GetGoods", err)
			}
			err = ValidateGetGoodsBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("shop", "GetGoods", err)
			}
			return nil, NewGetGoodsBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("shop", "GetGoods", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateGoodsRequest instantiates a HTTP request object with method and
// path set to call the "shop" service "UpdateGoods" endpoint
func (c *Client) BuildUpdateGoodsRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		goodsID string
	)
	{
		p, ok := v.(*shop.UpdateGoodsPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("shop", "UpdateGoods", "*shop.UpdateGoodsPayload", v)
		}
		goodsID = p.GoodsID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateGoodsShopPath(goodsID)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("shop", "UpdateGoods", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateGoodsRequest returns an encoder for requests sent to the shop
// UpdateGoods server.
func EncodeUpdateGoodsRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*shop.UpdateGoodsPayload)
		if !ok {
			return goahttp.ErrInvalidType("shop", "UpdateGoods", "*shop.UpdateGoodsPayload", v)
		}
		{
			head := p.JWTToken
			req.Header.Set("Authorization", head)
		}
		body := NewUpdateGoodsRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("shop", "UpdateGoods", err)
		}
		return nil
	}
}

// DecodeUpdateGoodsResponse returns a decoder for responses returned by the
// shop UpdateGoods endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeUpdateGoodsResponse may return the following errors:
//	- "internal_server_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeUpdateGoodsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateGoodsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("shop", "UpdateGoods", err)
			}
			p := NewUpdateGoodsGoodsSpuOK(&body)
			view := "default"
			vres := &shopviews.GoodsSpu{Projected: p, View: view}
			if err = shopviews.ValidateGoodsSpu(vres); err != nil {
				return nil, goahttp.ErrValidationError("shop", "UpdateGoods", err)
			}
			res := shop.NewGoodsSpu(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body UpdateGoodsInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("shop", "UpdateGoods", err)
			}
			err = ValidateUpdateGoodsInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("shop", "UpdateGoods", err)
			}
			return nil, NewUpdateGoodsInternalServerError(&body)
		case http.StatusBadRequest:
			var (
				body UpdateGoodsBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("shop", "UpdateGoods", err)
			}
			err = ValidateUpdateGoodsBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("shop", "UpdateGoods", err)
			}
			return nil, NewUpdateGoodsBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("shop", "UpdateGoods", resp.StatusCode, string(body))
		}
	}
}

// unmarshalGoodsResponseBodyToShopGoods builds a value of type *shop.Goods
// from a value of type *GoodsResponseBody.
func unmarshalGoodsResponseBodyToShopGoods(v *GoodsResponseBody) *shop.Goods {
	res := &shop.Goods{
		Errcode: *v.Errcode,
		Errmsg:  *v.Errmsg,
	}
	res.Data = make([]*shop.GoodsSpu, len(v.Data))
	for i, val := range v.Data {
		res.Data[i] = unmarshalGoodsSpuResponseBodyToShopGoodsSpu(val)
	}

	return res
}

// unmarshalGoodsSpuResponseBodyToShopGoodsSpu builds a value of type
// *shop.GoodsSpu from a value of type *GoodsSpuResponseBody.
func unmarshalGoodsSpuResponseBodyToShopGoodsSpu(v *GoodsSpuResponseBody) *shop.GoodsSpu {
	res := &shop.GoodsSpu{
		ID:          *v.ID,
		Name:        *v.Name,
		Description: *v.Description,
	}
	res.GoodsType = unmarshalGoodsTypeResponseBodyToShopGoodsType(v.GoodsType)
	res.GoodsBrand = unmarshalGoodsBrandResponseBodyToShopGoodsBrand(v.GoodsBrand)
	if v.GoodsSku != nil {
		res.GoodsSku = make([]*shop.GoodsSku, len(v.GoodsSku))
		for i, val := range v.GoodsSku {
			res.GoodsSku[i] = unmarshalGoodsSkuResponseBodyToShopGoodsSku(val)
		}
	}

	return res
}

// unmarshalGoodsTypeResponseBodyToShopGoodsType builds a value of type
// *shop.GoodsType from a value of type *GoodsTypeResponseBody.
func unmarshalGoodsTypeResponseBodyToShopGoodsType(v *GoodsTypeResponseBody) *shop.GoodsType {
	res := &shop.GoodsType{
		ID:          *v.ID,
		Name:        *v.Name,
		Description: *v.Description,
	}

	return res
}

// unmarshalGoodsBrandResponseBodyToShopGoodsBrand builds a value of type
// *shop.GoodsBrand from a value of type *GoodsBrandResponseBody.
func unmarshalGoodsBrandResponseBodyToShopGoodsBrand(v *GoodsBrandResponseBody) *shop.GoodsBrand {
	res := &shop.GoodsBrand{
		ID:          *v.ID,
		Name:        *v.Name,
		Description: *v.Description,
	}

	return res
}

// unmarshalGoodsSkuResponseBodyToShopGoodsSku builds a value of type
// *shop.GoodsSku from a value of type *GoodsSkuResponseBody.
func unmarshalGoodsSkuResponseBodyToShopGoodsSku(v *GoodsSkuResponseBody) *shop.GoodsSku {
	if v == nil {
		return nil
	}
	res := &shop.GoodsSku{
		ID:   *v.ID,
		Name: *v.Name,
	}
	if v.GoodsSkuValue != nil {
		res.GoodsSkuValue = make([]*shop.GoodsSkuValue, len(v.GoodsSkuValue))
		for i, val := range v.GoodsSkuValue {
			res.GoodsSkuValue[i] = unmarshalGoodsSkuValueResponseBodyToShopGoodsSkuValue(val)
		}
	}

	return res
}

// unmarshalGoodsSkuValueResponseBodyToShopGoodsSkuValue builds a value of type
// *shop.GoodsSkuValue from a value of type *GoodsSkuValueResponseBody.
func unmarshalGoodsSkuValueResponseBodyToShopGoodsSkuValue(v *GoodsSkuValueResponseBody) *shop.GoodsSkuValue {
	if v == nil {
		return nil
	}
	res := &shop.GoodsSkuValue{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// unmarshalGoodsTypeResponseBodyToShopviewsGoodsTypeView builds a value of
// type *shopviews.GoodsTypeView from a value of type *GoodsTypeResponseBody.
func unmarshalGoodsTypeResponseBodyToShopviewsGoodsTypeView(v *GoodsTypeResponseBody) *shopviews.GoodsTypeView {
	res := &shopviews.GoodsTypeView{
		ID:          v.ID,
		Name:        v.Name,
		Description: v.Description,
	}

	return res
}

// unmarshalGoodsBrandResponseBodyToShopviewsGoodsBrandView builds a value of
// type *shopviews.GoodsBrandView from a value of type *GoodsBrandResponseBody.
func unmarshalGoodsBrandResponseBodyToShopviewsGoodsBrandView(v *GoodsBrandResponseBody) *shopviews.GoodsBrandView {
	res := &shopviews.GoodsBrandView{
		ID:          v.ID,
		Name:        v.Name,
		Description: v.Description,
	}

	return res
}

// unmarshalGoodsSkuResponseBodyToShopviewsGoodsSkuView builds a value of type
// *shopviews.GoodsSkuView from a value of type *GoodsSkuResponseBody.
func unmarshalGoodsSkuResponseBodyToShopviewsGoodsSkuView(v *GoodsSkuResponseBody) *shopviews.GoodsSkuView {
	if v == nil {
		return nil
	}
	res := &shopviews.GoodsSkuView{
		ID:   v.ID,
		Name: v.Name,
	}
	if v.GoodsSkuValue != nil {
		res.GoodsSkuValue = make([]*shopviews.GoodsSkuValueView, len(v.GoodsSkuValue))
		for i, val := range v.GoodsSkuValue {
			res.GoodsSkuValue[i] = unmarshalGoodsSkuValueResponseBodyToShopviewsGoodsSkuValueView(val)
		}
	}

	return res
}

// unmarshalGoodsSkuValueResponseBodyToShopviewsGoodsSkuValueView builds a
// value of type *shopviews.GoodsSkuValueView from a value of type
// *GoodsSkuValueResponseBody.
func unmarshalGoodsSkuValueResponseBodyToShopviewsGoodsSkuValueView(v *GoodsSkuValueResponseBody) *shopviews.GoodsSkuValueView {
	if v == nil {
		return nil
	}
	res := &shopviews.GoodsSkuValueView{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}
// Code generated by goa v3.1.1, DO NOT EDIT.
//
// shop client
//
// Command:
// $ goa gen goa-shopping/design

package shop

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "shop" service client.
type Client struct {
	GoodsListEndpoint   goa.Endpoint
	GetGoodsEndpoint    goa.Endpoint
	UpdateGoodsEndpoint goa.Endpoint
}

// NewClient initializes a "shop" service client given the endpoints.
func NewClient(goodsList, getGoods, updateGoods goa.Endpoint) *Client {
	return &Client{
		GoodsListEndpoint:   goodsList,
		GetGoodsEndpoint:    getGoods,
		UpdateGoodsEndpoint: updateGoods,
	}
}

// GoodsList calls the "GoodsList" endpoint of the "shop" service.
func (c *Client) GoodsList(ctx context.Context, p *GoodsListPayload) (res *GoodsListResult, err error) {
	var ires interface{}
	ires, err = c.GoodsListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GoodsListResult), nil
}

// GetGoods calls the "GetGoods" endpoint of the "shop" service.
func (c *Client) GetGoods(ctx context.Context, p *GetGoodsPayload) (res *GoodsSpu, err error) {
	var ires interface{}
	ires, err = c.GetGoodsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GoodsSpu), nil
}

// UpdateGoods calls the "UpdateGoods" endpoint of the "shop" service.
func (c *Client) UpdateGoods(ctx context.Context, p *UpdateGoodsPayload) (res *GoodsSpu, err error) {
	var ires interface{}
	ires, err = c.UpdateGoodsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GoodsSpu), nil
}

// Code generated by goa v3.1.1, DO NOT EDIT.
//
// shop HTTP client CLI support package
//
// Command:
// $ goa gen goa-shopping/design

package client

import (
	"encoding/json"
	"fmt"
	shop "goa-shopping/gen/shop"
	"strconv"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// BuildGoodsListPayload builds the payload for the shop GoodsList endpoint
// from CLI flags.
func BuildGoodsListPayload(shopGoodsListCursor string, shopGoodsListLimit string, shopGoodsListJWTToken string) (*shop.GoodsListPayload, error) {
	var err error
	var cursor int
	{
		var v int64
		v, err = strconv.ParseInt(shopGoodsListCursor, 10, 64)
		cursor = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for cursor, must be INT")
		}
	}
	var limit int
	{
		var v int64
		v, err = strconv.ParseInt(shopGoodsListLimit, 10, 64)
		limit = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for limit, must be INT")
		}
	}
	var jwtToken *string
	{
		if shopGoodsListJWTToken != "" {
			jwtToken = &shopGoodsListJWTToken
		}
	}
	v := &shop.GoodsListPayload{}
	v.Cursor = cursor
	v.Limit = limit
	v.JWTToken = jwtToken

	return v, nil
}

// BuildGetGoodsPayload builds the payload for the shop GetGoods endpoint from
// CLI flags.
func BuildGetGoodsPayload(shopGetGoodsGoodsID string, shopGetGoodsJWTToken string) (*shop.GetGoodsPayload, error) {
	var err error
	var goodsID string
	{
		goodsID = shopGetGoodsGoodsID
		if utf8.RuneCountInString(goodsID) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("goodsID", goodsID, utf8.RuneCountInString(goodsID), 1, true))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwtToken string
	{
		jwtToken = shopGetGoodsJWTToken
	}
	v := &shop.GetGoodsPayload{}
	v.GoodsID = goodsID
	v.JWTToken = jwtToken

	return v, nil
}

// BuildUpdateGoodsPayload builds the payload for the shop UpdateGoods endpoint
// from CLI flags.
func BuildUpdateGoodsPayload(shopUpdateGoodsBody string, shopUpdateGoodsGoodsID string, shopUpdateGoodsJWTToken string) (*shop.UpdateGoodsPayload, error) {
	var err error
	var body UpdateGoodsRequestBody
	{
		err = json.Unmarshal([]byte(shopUpdateGoodsBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"description\": \"水果x刘海屏就很丑\",\n      \"name\": \"水果x\"\n   }'")
		}
		if utf8.RuneCountInString(body.Name) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 20, false))
		}
		if utf8.RuneCountInString(body.Description) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", body.Description, utf8.RuneCountInString(body.Description), 128, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var goodsID string
	{
		goodsID = shopUpdateGoodsGoodsID
		if utf8.RuneCountInString(goodsID) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("goodsID", goodsID, utf8.RuneCountInString(goodsID), 1, true))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwtToken string
	{
		jwtToken = shopUpdateGoodsJWTToken
	}
	v := &shop.UpdateGoodsPayload{
		Name:        body.Name,
		Description: body.Description,
	}
	v.GoodsID = goodsID
	v.JWTToken = jwtToken

	return v, nil
}

package controller

import (
	"context"
	"fmt"
	log "goa-shopping/gen/log"
	shop "goa-shopping/gen/shop"

	"goa.design/goa/v3/security"
)

// shop service example implementation.
// The example methods log the requests and return zero values.
type shopsrvc struct {
	logger *log.Logger
}

// NewShop returns the shop service implementation.
func NewShop(logger *log.Logger) shop.Service {
	return &shopsrvc{logger}
}

// JWTAuth implements the authorization logic for service "shop" for the "jwt"
// security scheme.
func (s *shopsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
}

// 商品列表
func (s *shopsrvc) GoodsList(ctx context.Context, p *shop.GoodsListPayload) (res *shop.GoodsListResult, err error) {
	res = &shop.GoodsListResult{}
	s.logger.Info("shop.GoodsList")
	return
}

// 商品详情
func (s *shopsrvc) GetGoods(ctx context.Context, p *shop.GetGoodsPayload) (res *shop.GoodsSpu, err error) {
	res = &shop.GoodsSpu{}
	s.logger.Info("shop.GetGoods")
	return
}

// 更新商品
func (s *shopsrvc) UpdateGoods(ctx context.Context, p *shop.UpdateGoodsPayload) (res *shop.GoodsSpu, err error) {
	res = &shop.GoodsSpu{}
	s.logger.Info("shop.UpdateGoods")
	return
}

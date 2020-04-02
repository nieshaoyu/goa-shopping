package dao

import (
	"context"

	"github.com/jinzhu/gorm"

	"goa-shopping/gen/auth"
	"goa-shopping/gen/log"
	"goa-shopping/model"
)

type UserDaoImpl interface {
	CreateUser(p *auth.SignUpPayload) (model.User, error)
}

type userDao struct {
	db     *gorm.DB
	logger *log.Logger
	ctx    context.Context
}

func NewUserDao(ctx context.Context, db *gorm.DB, logger *log.Logger) UserDaoImpl {
	return &userDao{
		db:     db,
		logger: logger,
		ctx:    ctx,
	}
}

func (u *userDao) CreateUser(p *auth.SignUpPayload) (model.User, error) {
	panic("impl me")
}

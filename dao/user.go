package dao

import (
	"context"
	"errors"
	"time"

	"github.com/jinzhu/gorm"

	"goa-shopping/gen/auth"
	"goa-shopping/gen/log"
	"goa-shopping/model"
)

type UserDaoImpl interface {
	// 创建用户
	CreateUser(p *auth.SignUpPayload) (model.User, error)
	// 查找用户
	FindUserByMobileOrEmail(mobile string) (model.User, error)
	// 更新登录时间
	UpdateLoginTime(user *model.User) error
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
	if p.Email == nil && p.Phone == nil {
		u.logger.Info("手机号和邮箱不能同时为空")
		return model.User{}, errors.New("手机号和邮箱不能同时为空")
	}
	var (
		mobile,
		email string
	)
	if p.Email != nil && *p.Email != "" {
		email = *p.Email
	}
	if p.Phone != nil && *p.Phone != "" {
		mobile = *p.Phone
	}
	user := model.User{
		Username:  p.Nickname,
		Mobile:    mobile,
		Email:     email,
		Nickname:  p.Nickname,
		Type:      model.UserTypeRegular,
		IsActive:  true,
		LoginTime: nil,
	}
	// 没有配置邮箱, 没做VerifyCode验证
	pwd, err := user.CreatePassword(p.Password)
	if err != nil {
		u.logger.Error("加密失败")
		return model.User{}, errors.New("创建用户失败")
	}
	user.Password = pwd
	err = u.db.Model(user).Create(&user).Error
	return user, err
}

func (u *userDao) FindUserByMobileOrEmail(mobile string) (model.User, error) {
	var user model.User
	// gorm 已经做了逻辑删除校验
	err := u.db.Model(model.User{}).Where("mobile = ? or email = ?", mobile, mobile).Find(&user).Error
	return user, err
}

// 更新登陆时间
func (u *userDao) UpdateLoginTime(user *model.User) error {
	if err := u.db.Model(user).UpdateColumn("login_time", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

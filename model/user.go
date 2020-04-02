package model

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

const (
	UserTypeRegular int = 1
	UserTypeAdmin       = 2 // 管理员
)

type User struct {
	BaseModel

	Username string `gorm:"type:varchar(36);not null;unique;"`
	Password string `gorm:"type:varchar(128);"`
	Mobile   string `gorm:"type:varchar(15);not null;unique;"`
	Email    string `gorm:"type:varchar(20);not null;unique;"`
	Nickname string `gorm:"type:varchar(36);not null;default:''"`
	// 用户类型
	Type int `gorm:"type:int;not null;default:1"`

	IsActive bool `gorm:"type:tinyint;not null;default:true"`
	// 登录时间
	LoginTime *time.Time
}

// 账户登录日志
type UserLoginLog struct {
	// 账户登录日志ID
	ID int `gorm:"primary_key"`

	// 用户ID，
	UserID string `gorm:"type:varchar(36);index:idx_login_log_user_id"`

	// 用户登录的ip
	IPAddr string `gorm:"type:varchar(64)"`

	UserAgent string `gorm:"type:text"`
	CreateAt  *time.Time
}

// CreatePassword 加密密码
func (u *User) CreatePassword(raw string) (string, error) {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the DefaultCost (10)

	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash), nil
}

// CheckPassword 检查密码
func (u *User) CheckPassword(plainPwd, hashedPwd string) bool {

	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		return false
	}

	return true
}

// Scopes
func (u *User) Scopes() []string {
	if u.Type == UserTypeRegular {
		return []string{"role:user"}
	} else if u.Type == UserTypeAdmin {
		return []string{"role:admin"}
	}
	return make([]string, 0, 0)
}

// 是否重复
func IsDuplicateError(err error) bool {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if ok {
		if mysqlErr.Number == 1062 {
			return true
		}
	}

	return false
}

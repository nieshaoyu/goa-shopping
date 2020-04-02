package dao

import (
	"time"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"

	"goa-shopping/config"
	"goa-shopping/model"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) {
	zap.L().Debug("connect db", zap.String("dsn", cfg.Database.DSN))

	var err error

	DB, err = gorm.Open("mysql", cfg.Database.DSN)
	if err != nil {
		zap.L().Panic("connect db failed", zap.Error(err))
	}

	if cfg.Debug {
		DB.LogMode(cfg.Database.LogMode) // 是否打印SQL
	}

	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	if cfg.Database.MaxIdleConns > 0 {
		DB.DB().SetMaxIdleConns(cfg.Database.MaxIdleConns)
	}

	// SetMaxOpenCons 设置数据库的最大连接数量。
	if cfg.Database.MaxOpenConns > 0 {
		DB.DB().SetMaxOpenConns(cfg.Database.MaxOpenConns)
	}

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	if cfg.Database.ConnMaxLifetime != "" {
		maxLifetime, err := time.ParseDuration(cfg.Database.ConnMaxLifetime)
		if err != nil {
			zap.L().Panic("db ConnMaxLifetime parse failed", zap.Error(err))
		}

		DB.DB().SetConnMaxLifetime(maxLifetime)
	}

	if err := DB.DB().Ping(); err != nil {
		zap.L().Panic("ping db failed", zap.Error(err))
	}
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

func AutoMigrateDB() {
	if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci").
		AutoMigrate(
			&model.GoodsBrand{},
			&model.User{},
			&model.UserLoginLog{},
			&model.GoodsSKU{},
			&model.GoodsSPU{},
			&model.GoodsSPUValue{},
			&model.GoodsType{},
		).Error; err != nil {
		zap.L().Panic("migrate db fail", zap.Error(err))
	}
}

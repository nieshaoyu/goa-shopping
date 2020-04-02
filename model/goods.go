package model

// 商品类型
type GoodsType struct {
	BaseModel
	Name        string `gorm:"type:varchar(36);not null;unique;"`
	Description string `gorm:"type:varchar(50);not null;"`
}

// 商品品牌
type GoodsBrand struct {
	BaseModel
	Name        string `gorm:"type:varchar(36);not null;unique;"`
	Description string `gorm:"type:varchar(50);not null;"`
}

type GoodsSPU struct {
	BaseModel
	// 类型ID
	TypeID int `gorm:"type:int;not null;"`
	// 品牌ID
	BrandID int `gorm:"type:int;not null;"`

	// SKU  ID
	SkuID int `gorm:"type:int;not null;"`
}

// 商品SKU
type GoodsSKU struct {
	BaseModel
	Name string `gorm:"type:varchar(36);not null;unique;"`
}

type GoodsSPUValue struct {
	BaseModel
	Name string `gorm:"type:varchar(36);not null;unique;"`
	// SKU ID
	SkuID int `gorm:"type:int;not null;"`
}

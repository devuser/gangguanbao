package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

// Material 材质
type Material struct {
	ID        int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name      string     `gorm:"column:name;not null" json:"name"`
	Code      string     `gorm:"column:code" json:"code"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (Material) TableName() string { return "biz_materials" }

// Spec 规格
type Spec struct {
	ID           int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name         string     `gorm:"column:name;not null" json:"name"`
	MaterialID   *int       `gorm:"column:material_id" json:"materialId"`
	MaterialName string     `gorm:"-" json:"materialName"`
	Unit         string     `gorm:"column:unit;default:件" json:"unit"`
	CreatedAt    *time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (Spec) TableName() string { return "biz_specs" }

// Vendor 厂家
type Vendor struct {
	ID        int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name      string     `gorm:"column:name;not null" json:"name"`
	Code      string     `gorm:"column:code" json:"code"`
	Contact   string     `gorm:"column:contact" json:"contact"`
	Phone     string     `gorm:"column:phone" json:"phone"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (Vendor) TableName() string { return "biz_vendors" }

// Order 订单
type Order struct {
	ID         int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrderNo    string     `gorm:"column:order_no" json:"orderNo"`
	MaterialID *int       `gorm:"column:material_id" json:"materialId"`
	SpecID     *int       `gorm:"column:spec_id" json:"specId"`
	VendorID   *int       `gorm:"column:vendor_id" json:"vendorId"`
	OrderDate  *string    `gorm:"column:order_date" json:"orderDate"`
	Quantity   float64    `gorm:"column:quantity;default:0" json:"quantity"`
	UnitPrice  *float64   `gorm:"column:unit_price" json:"unitPrice"`
	Amount     *float64   `gorm:"column:amount" json:"amount"`
	Remark     string     `gorm:"column:remark" json:"remark"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"updatedAt"`
	// 关联名称（JOIN 查询填充，只读不写入）
	MaterialName string `gorm:"column:material_name;->" json:"materialName"`
	SpecName     string `gorm:"column:spec_name;->" json:"specName"`
	VendorName   string `gorm:"column:vendor_name;->" json:"vendorName"`
}

func (Order) TableName() string { return "biz_orders" }

package repository

import (
	"fmt"
	"strings"

	"github.com/boyosoft/gangguanbao/backend/internal/model"
	"gorm.io/gorm"
)

// OrderQuery 订单查询条件
type OrderQuery struct {
	MaterialID *int   `form:"materialId"`
	SpecID     *int   `form:"specId"`
	VendorID   *int   `form:"vendorId"`
	OrderNo    string `form:"orderNo"`
	DateFrom   string `form:"dateFrom"`
	DateTo     string `form:"dateTo"`
	Page       int    `form:"page"`
	PageSize   int    `form:"pageSize"`
	SortBy     string `form:"sortBy"`
	SortOrder  string `form:"sortOrder"` // asc / desc
}

// ListOrders 分页查询订单，支持多条件搜索和排序
func (r *Repository) ListOrders(q OrderQuery) ([]model.Order, int64, error) {
	query := r.DB.Table("biz_orders o").
		Select(`o.id, o.order_no, o.material_id, o.spec_id, o.vendor_id,
			o.order_date, o.quantity, o.unit_price, o.amount, o.remark,
			o.created_at, o.updated_at,
			COALESCE(m.name, '') AS material_name,
			COALESCE(s.name, '') AS spec_name,
			COALESCE(v.name, '') AS vendor_name`).
		Joins("LEFT JOIN biz_materials m ON o.material_id = m.id").
		Joins("LEFT JOIN biz_specs s ON o.spec_id = s.id").
		Joins("LEFT JOIN biz_vendors v ON o.vendor_id = v.id")

	query = applyOrderWhere(query, q)

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	orderClause := "o.id DESC"
	if q.SortBy != "" {
		sortCol := map[string]string{
			"orderDate":    "o.order_date",
			"orderNo":      "o.order_no",
			"quantity":     "o.quantity",
			"unitPrice":    "o.unit_price",
			"amount":       "o.amount",
			"createdAt":    "o.created_at",
			"materialName": "material_name",
			"specName":     "spec_name",
			"vendorName":   "vendor_name",
		}
		if col, ok := sortCol[q.SortBy]; ok {
			ord := "ASC"
			if strings.ToLower(q.SortOrder) == "desc" {
				ord = "DESC"
			}
			orderClause = fmt.Sprintf("%s %s", col, ord)
		}
	}

	page := q.Page
	if page < 1 {
		page = 1
	}
	pageSize := q.PageSize
	if pageSize < 1 || pageSize > 500 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	var list []model.Order
	if err := query.Order(orderClause).Offset(offset).Limit(pageSize).Scan(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// ListOrdersForExport 查询订单（用于导出，无分页限制，最多 10000 条）
func (r *Repository) ListOrdersForExport(q OrderQuery) ([]model.Order, error) {
	query := r.DB.Table("biz_orders o").
		Select(`o.id, o.order_no, o.material_id, o.spec_id, o.vendor_id,
			o.order_date, o.quantity, o.unit_price, o.amount, o.remark,
			o.created_at, o.updated_at,
			COALESCE(m.name, '') AS material_name,
			COALESCE(s.name, '') AS spec_name,
			COALESCE(v.name, '') AS vendor_name`).
		Joins("LEFT JOIN biz_materials m ON o.material_id = m.id").
		Joins("LEFT JOIN biz_specs s ON o.spec_id = s.id").
		Joins("LEFT JOIN biz_vendors v ON o.vendor_id = v.id")

	query = applyOrderWhere(query, q)

	orderClause := "o.id DESC"
	if q.SortBy != "" {
		sortCol := map[string]string{
			"orderDate": "o.order_date", "orderNo": "o.order_no", "quantity": "o.quantity",
			"unitPrice": "o.unit_price", "amount": "o.amount", "createdAt": "o.created_at",
			"materialName": "material_name", "specName": "spec_name", "vendorName": "vendor_name",
		}
		if col, ok := sortCol[q.SortBy]; ok {
			ord := "ASC"
			if strings.ToLower(q.SortOrder) == "desc" {
				ord = "DESC"
			}
			orderClause = fmt.Sprintf("%s %s", col, ord)
		}
	}

	var list []model.Order
	err := query.Order(orderClause).Limit(10000).Scan(&list).Error
	return list, err
}

// GetOrder 根据ID获取订单（含关联名称）
func (r *Repository) GetOrder(id int) (*model.Order, error) {
	var o model.Order
	err := r.DB.Table("biz_orders o").
		Select(`o.id, o.order_no, o.material_id, o.spec_id, o.vendor_id,
			o.order_date, o.quantity, o.unit_price, o.amount, o.remark,
			o.created_at, o.updated_at,
			COALESCE(m.name, '') AS material_name,
			COALESCE(s.name, '') AS spec_name,
			COALESCE(v.name, '') AS vendor_name`).
		Joins("LEFT JOIN biz_materials m ON o.material_id = m.id").
		Joins("LEFT JOIN biz_specs s ON o.spec_id = s.id").
		Joins("LEFT JOIN biz_vendors v ON o.vendor_id = v.id").
		Where("o.id = ?", id).
		First(&o).Error
	if err != nil {
		return nil, err
	}
	return &o, nil
}

// CreateOrder 新增订单
func (r *Repository) CreateOrder(o *model.Order) error {
	return r.DB.Create(o).Error
}

// UpdateOrder 更新订单
func (r *Repository) UpdateOrder(o *model.Order) error {
	return r.DB.Save(o).Error
}

// DeleteOrder 删除订单
func (r *Repository) DeleteOrder(id int) error {
	return r.DB.Delete(&model.Order{}, id).Error
}

func applyOrderWhere(query *gorm.DB, q OrderQuery) *gorm.DB {
	db := query
	if q.MaterialID != nil && *q.MaterialID > 0 {
		db = db.Where("o.material_id = ?", *q.MaterialID)
	}
	if q.SpecID != nil && *q.SpecID > 0 {
		db = db.Where("o.spec_id = ?", *q.SpecID)
	}
	if q.VendorID != nil && *q.VendorID > 0 {
		db = db.Where("o.vendor_id = ?", *q.VendorID)
	}
	if q.OrderNo != "" {
		db = db.Where("o.order_no LIKE ?", "%"+q.OrderNo+"%")
	}
	if q.DateFrom != "" {
		db = db.Where("o.order_date >= ?", q.DateFrom)
	}
	if q.DateTo != "" {
		db = db.Where("o.order_date <= ?", q.DateTo)
	}
	return db
}

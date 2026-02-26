package repository

import (
	"time"
	// "gorm.io/gorm"
)

// StatsCounts 基础数据条数
type StatsCounts struct {
	OrderCount    int64 `json:"orderCount"`
	MaterialCount int64 `json:"materialCount"`
	SpecCount     int64 `json:"specCount"`
	VendorCount   int64 `json:"vendorCount"`
}

// YearCount 年度订单数
type YearCount struct {
	Year  int   `json:"year"`
	Count int64 `json:"count"`
}

// MonthCount 月度订单数
type MonthCount struct {
	Month int   `json:"month"`
	Count int64 `json:"count"`
}

// NameCount 按名称统计（材质/规格）
type NameCount struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

// GetStatsCounts 获取各表条数
func (r *Repository) GetStatsCounts() (StatsCounts, error) {
	var s StatsCounts
	if err := r.DB.Table("biz_orders").Count(&s.OrderCount).Error; err != nil {
		return s, err
	}
	if err := r.DB.Table("biz_materials").Count(&s.MaterialCount).Error; err != nil {
		return s, err
	}
	if err := r.DB.Table("biz_specs").Count(&s.SpecCount).Error; err != nil {
		return s, err
	}
	if err := r.DB.Table("biz_vendors").Count(&s.VendorCount).Error; err != nil {
		return s, err
	}
	return s, nil
}

// GetOrdersByYear 每年订单数
func (r *Repository) GetOrdersByYear() ([]YearCount, error) {
	var list []YearCount
	err := r.DB.Table("biz_orders").
		Select("YEAR(order_date) AS year, COUNT(*) AS count").
		Where("order_date IS NOT NULL").
		Group("YEAR(order_date)").
		Order("year ASC").
		Scan(&list).Error
	return list, err
}

// GetOrdersByMonth 指定年份各月订单数（1-12，无数据月份为0）
func (r *Repository) GetOrdersByMonth(year int) ([]MonthCount, error) {
	result := make([]MonthCount, 12)
	for i := 1; i <= 12; i++ {
		result[i-1] = MonthCount{Month: i, Count: 0}
	}

	var list []MonthCount
	err := r.DB.Table("biz_orders").
		Select("MONTH(order_date) AS month, COUNT(*) AS count").
		Where("order_date IS NOT NULL AND YEAR(order_date) = ?", year).
		Group("MONTH(order_date)").
		Order("month ASC").
		Scan(&list).Error
	if err != nil {
		return nil, err
	}
	for _, m := range list {
		if m.Month >= 1 && m.Month <= 12 {
			result[m.Month-1] = m
		}
	}
	return result, nil
}

// GetOrdersByMaterial 指定年份各材质订单数
func (r *Repository) GetOrdersByMaterial(year int) ([]NameCount, error) {
	var list []NameCount
	err := r.DB.Table("biz_orders o").
		Select("COALESCE(m.name, '未分类') AS name, COUNT(*) AS count").
		Joins("LEFT JOIN biz_materials m ON o.material_id = m.id").
		Where("YEAR(o.order_date) = ?", year).
		Group("o.material_id, m.name").
		Order("count DESC").
		Scan(&list).Error
	return list, err
}

// GetOrdersBySpec 指定年份各规格订单数
func (r *Repository) GetOrdersBySpec(year int) ([]NameCount, error) {
	var list []NameCount
	err := r.DB.Table("biz_orders o").
		Select("COALESCE(s.name, '未分类') AS name, COUNT(*) AS count").
		Joins("LEFT JOIN biz_specs s ON o.spec_id = s.id").
		Where("YEAR(o.order_date) = ?", year).
		Group("o.spec_id, s.name").
		Order("count DESC").
		Scan(&list).Error
	return list, err
}

// GetCurrentYear 当前年份
func GetCurrentYear() int {
	return time.Now().Year()
}
